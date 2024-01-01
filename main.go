package main

import (
	"context"
	"github.com/chenbihao/gob/provider/demo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chenbihao/gob/framework/gin"
)

func main() {

	// 核心框架初始化
	// core := framework.NewCore()
	core := gin.New()

	// 绑定具体的服务
	core.Bind(&demo.DemoServiceProvider{})

	// 设置路由
	registerRouter(core)

	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8080",
	}

	// 这个goroutine是启动服务的goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 当前的 goroutine 等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 设置超时关闭限制
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 调用 Server.Shutdown graceful 结束
	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

}
