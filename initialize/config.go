package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go-Framework/global"
)

func InitializeConfig() {
	// 设置配置文件路径
	configPath := "config.yaml"
	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("InitializeConfig -- 读取配置文件错误: %s \n", err.Error()))
	}
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Config); err != nil {
		panic(fmt.Errorf("InitializeConfig -- 写入配置文件错误: %s \n", err.Error()))
	}
}
