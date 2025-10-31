package middleware

import (
	"bytes"
	"context"
	"io"
	"time"

	"go-Framework/global"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// responseBodyWriter 用于捕获响应体
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logger 日志中间件 - 记录请求和响应信息
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成唯一的 request_id
		requestID := uuid.New().String()

		// 将 request_id 存储到 context 中
		ctx := context.WithValue(c.Request.Context(), "request_id", requestID)
		c.Request = c.Request.WithContext(ctx)

		// 记录请求开始时间
		startTime := time.Now()

		// 读取请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 重新设置请求体，因为 Body 只能读取一次
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 创建自定义的 ResponseWriter 以捕获响应体
		responseWriter := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = responseWriter

		// 处理请求
		c.Next()

		// 计算请求耗时
		duration := time.Since(startTime)

		// 获取响应体
		responseBody := responseWriter.body.String()

		// 打印一条完整的日志，包含请求和响应信息
		global.App.Log.Info().
			Str("request_id", requestID).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("query", c.Request.URL.RawQuery).
			Str("client_ip", c.ClientIP()).
			Str("user_agent", c.Request.UserAgent()).
			Str("request_body", string(requestBody)).
			Int("status_code", c.Writer.Status()).
			Str("response_body", responseBody).
			Dur("duration", duration).
			Str("duration_ms", duration.String()).
			Msg("请求处理完成")
	}
}
