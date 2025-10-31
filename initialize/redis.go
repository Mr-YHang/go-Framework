package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-Framework/global"
)

func InitializeRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", global.App.Config.Redis.Host, global.App.Config.Redis.Port),
		Password: global.App.Config.Redis.Password, // no password set
		DB:       global.App.Config.Redis.DB,       // use default DB
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Errorf("InitializeRedis -- 初始化redis错误 : %s \n", err.Error()))
	}

	global.App.Redis = client
}
