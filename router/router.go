package router

import (
	"github.com/gin-gonic/gin"
	"go-Framework/app/handler"
	"net/http"
)

func Router(router *gin.Engine, h handler.Handler) {
	// 健康检查
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	user := router.Group("/user")
	user.POST("/login", h.Session.Login)

}
