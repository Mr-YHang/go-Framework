package config

type Config struct {
	App     App       `mapstructure:"app" json:"app" yaml:"app"`
	Log     Log       `mapstructure:"log" json:"log" yaml:"log"`
	OrderDB DBConf    `mapstructure:"order_db" json:"order_db" yaml:"order_db"`
	Redis   RedisConf `mapstructure:"redis" json:"redis" yaml:"redis"`
}

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`                // 环境
	Port    string `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"` // 应用名
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`    // 域名
}

type Log struct {
	IsConsole  bool   `mapstructure:"is_console" json:"is_console" yaml:"is_console"`    // 命令台显示
	Path       string `mapstructure:"path" json:"path" yaml:"path"`                      // 日志路径
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`             // 保留天数
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`          // 单文件最大M
	Level      int    `mapstructure:"level" json:"level" yaml:"level"`                   // 等级
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"` // 最大文件数
}

type DBConf struct {
	Driver              string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Host                string `mapstructure:"host" json:"host" yaml:"host"`
	Port                int    `mapstructure:"port" json:"port" yaml:"port"`
	Database            string `mapstructure:"database" json:"database" yaml:"database"`
	UserName            string `mapstructure:"username" json:"username" yaml:"username"`
	Password            string `mapstructure:"password" json:"password" yaml:"password"`
	Charset             string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdleConns        int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns        int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`
}

type RedisConf struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
