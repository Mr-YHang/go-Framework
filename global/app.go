package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"go-Framework/config"
	"gorm.io/gorm"
)

type Application struct {
	Config *config.Config // 全局配置文件
	Log    zerolog.Logger // 全局日志
	DB     *gorm.DB       // 全局mysql
	Redis  *redis.Client  // 全局redis
}

var App = new(Application)
