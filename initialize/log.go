package initialize

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"go-Framework/config"
	"go-Framework/global"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"runtime/debug"
	"time"
)

func InitializeLog(config config.Config) {
	// 设置错误堆栈的 marshaler 为 pkgerrors.MarshalStack
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	// 设置时间字段的格式为 RFC3339Nano
	zerolog.TimeFieldFormat = time.RFC3339Nano

	// 将配置中的日志级别转换为 zerolog.Level 类型
	var logLevel = zerolog.Level(config.Log.Level)
	// 检查配置的日志级别是否在有效范围内，如果不在则默认设置为 INFO 级别
	if logLevel < -1 || logLevel > 7 {
		logLevel = zerolog.InfoLevel
	}

	// 定义一个切片，用于存储日志写入器
	var writers []io.Writer

	// 如果启用了控制台日志输出
	if config.Log.IsConsole {
		// 向 writers 切片中添加一个控制台写入器
		writers = append(writers, zerolog.ConsoleWriter{
			Out:        os.Stderr,    // 输出到标准错误输出
			TimeFormat: time.RFC3339, // 设置时间格式
			FieldsExclude: []string{
				"user_agent",   // 排除 user_agent 字段
				"git_revision", // 排除 git_revision 字段
				"go_version",   // 排除 go_version 字段
			},
		})
	}

	// 如果启用了文件日志输出
	if len(config.Log.Path) > 0 {
		// 向 writers 切片中添加一个滚动文件写入器
		writers = append(writers, newRollingFile(config))
	}

	// 创建一个多写入器，将所有的写入器组合在一起
	mw := io.MultiWriter(writers...)

	// 定义一个变量，用于存储 git 版本信息
	var gitRevision string

	// 读取构建信息
	buildInfo, ok := debug.ReadBuildInfo()

	// 如果成功读取构建信息
	if ok {
		// 遍历构建信息中的设置项
		for _, v := range buildInfo.Settings {
			// 如果找到 vcs.revision 项，则将其值赋给 gitRevision 变量
			if v.Key == "vcs.revision" {
				gitRevision = v.Value
				break
			}
		}
	}

	// 初始化全局日志记录器
	global.App.Log = zerolog.New(mw).
		Level(zerolog.Level(logLevel)). // 设置日志级别
		With().
		Str("git_revision", gitRevision). // 添加 git 版本信息
		Str("go_version", buildInfo.GoVersion). // 添加 Go 版本信息
		Timestamp(). // 添加时间戳
		Logger()
}

// 创建一个滚动文件写入器
func newRollingFile(config config.Config) io.Writer {
	// 创建日志目录，如果目录已存在则不会报错
	if err := os.MkdirAll(config.Log.Path, 0744); err != nil {
		panic(fmt.Errorf("newRollingFile -- 创建日志文件路径错误: %s \n", err.Error()))
	}
	// 创建日志文件名 , 项目名+时间
	filename := config.App.AppName + "_" + time.Now().Format("2006-01-02") + ".log"

	// 返回一个 lumberjack.Logger 实例，用于处理日志文件的滚动
	return &lumberjack.Logger{
		Filename:   path.Join(config.Log.Path, "/", filename), // 日志文件路径
		MaxBackups: config.Log.MaxBackups,                     // 保留的滚动日志文件数量
		MaxSize:    config.Log.MaxSize,                        // 日志文件最大大小
		MaxAge:     config.Log.MaxAge,                         // 日志文件保留天数
	}
}
