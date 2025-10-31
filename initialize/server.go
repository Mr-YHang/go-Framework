package initialize

import (
	"context"
	"go-Framework/app/dao"
	"go-Framework/app/handler"
	"go-Framework/app/services"
	"go-Framework/global"
	"go-Framework/middleware"
	"go-Framework/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	
	// 注册全局中间件
	r.Use(middleware.Recover())
	
	// 初始化dao
	appDap := dao.NewDao()
	// 初始化服务层
	appServices := services.NewServices(appDap)
	// 初始化控制器层
	appHandler := handler.NewHandler(appServices)

	router.Router(r, *appHandler)

	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务启动监听失败--listen: %s\n", err.Error())
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务 ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("服务器等待关闭时发生错误:%s", err.Error())
	}
	log.Println("服务器已关闭")
}
