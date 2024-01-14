package command

import (
	"context"
	"fmt"
	"github.com/chenbihao/gob/framework/cobra"
	"github.com/chenbihao/gob/framework/contract"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
	命令介绍：
		web app 业务应用控制命令
	前置需求：
	支持命令：
		./gob app start --address=:8080
	支持配置：
		```app.yaml

		```
**/

// app启动地址
var appAddress = ""

// initAppCommand 初始化app命令和其子命令
func initAppCommand() *cobra.Command {
	// 设置启动地址
	appStartCommand.Flags().StringVar(&appAddress, "address", "8080", "设置app启动的地址，默认为8080端口")

	appCommand.AddCommand(appStartCommand)
	return appCommand
}

// AppCommand 是命令行参数第一级为app的命令，它没有实际功能，只是打印帮助文档
var appCommand = &cobra.Command{
	Use:   "app",
	Short: "业务应用控制命令",
	Long:  "业务应用控制命令，其包含业务启动，关闭，重启，查询等功能",
	RunE: func(c *cobra.Command, args []string) error {
		// 打印帮助文档
		c.Help()
		return nil
	},
}

// appStartCommand 启动一个Web服务
var appStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动一个Web服务",
	RunE: func(c *cobra.Command, args []string) error {
		// 从Command中获取服务容器
		container := c.GetContainer()
		// 从服务容器中获取kernel的服务实例
		kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
		// 从kernel服务实例中获取引擎
		core := kernelService.HttpEngine()

		appAddress = fmt.Sprintf(":" + appAddress)
		// 创建一个Server服务
		server := &http.Server{
			Handler: core,
			Addr:    appAddress,
		}
		// 这个goroutine是启动服务的goroutine
		go func() {
			server.ListenAndServe()
		}()
		log.Println("Server Started , Local: http://localhost" + appAddress)

		// 当前的goroutine等待信号量
		quit := make(chan os.Signal)
		// 监控信号：SIGINT, SIGTERM, SIGQUIT
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		// 这里会阻塞当前goroutine等待信号
		<-quit

		// 调用Server.Shutdown graceful结束
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}

		return nil

	},
}
