package global

import (
	"context"

	"go-Framework/config"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Application struct {
	Config *config.Config // 全局配置文件
	Log    zerolog.Logger // 全局日志
	DB     *gorm.DB       // 全局mysql
	Redis  *redis.Client  // 全局redis
}

var App = new(Application)

// LogWithContext 返回一个带有 context 信息的 logger（自动包含 request_id）
func LogWithContext(ctx context.Context) *zerolog.Logger {
	logger := App.Log.With().Logger()

	// 从 context 中提取 request_id
	if requestID, ok := ctx.Value("request_id").(string); ok {
		logger = logger.With().Str("request_id", requestID).Logger()
	}

	return &logger
}
