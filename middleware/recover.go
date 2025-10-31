package middleware

import (
	"github.com/gin-gonic/gin"
	"go-Framework/app/resp"
	"go-Framework/global"
	"runtime/debug"
)

// Recover 捕获panic并恢复，防止程序崩溃
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 获取堆栈信息
				stack := string(debug.Stack())

				// 记录panic日志
				global.App.Log.Error().
					Str("path", c.Request.URL.Path).
					Str("method", c.Request.Method).
					Str("client_ip", c.ClientIP()).
					Interface("error", err).
					Str("stack", stack).
					Msg("发生panic异常")

				resp.Fail(c, global.ProcessErrCode, "panic -- 服务器内部错误")

				// 终止后续处理
				c.Abort()
			}
		}()

		// 继续处理请求
		c.Next()
	}
}
