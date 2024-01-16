package main

import (
	"github.com/chenbihao/gob/app/console"
	"github.com/chenbihao/gob/app/http"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/provider/app"
	"github.com/chenbihao/gob/framework/provider/config"
	"github.com/chenbihao/gob/framework/provider/distributed"
	"github.com/chenbihao/gob/framework/provider/env"
	"github.com/chenbihao/gob/framework/provider/id"
	"github.com/chenbihao/gob/framework/provider/kernel"
	"github.com/chenbihao/gob/framework/provider/log"
	"github.com/chenbihao/gob/framework/provider/trace"
	"os"
)

func main() {

	// 如果仅当做脚手架来用的话，仅挂载new命令，不绑定各类服务提供者（暂时解决构建后无配置目录问题）
	args := os.Args
	if len(args) > 1 && args[1] == "new" {
		container := framework.NewGobContainer()
		console.RunRootCommand(container, true)
		return
	}

	// 初始化服务容器
	container := framework.NewGobContainer()
	// 绑定 App 服务提供者
	container.Bind(&app.AppProvider{})
	// 绑定 环境变量 服务提供者
	container.Bind(&env.EnvProvider{})
	// 绑定 分布式本地锁 服务提供者
	container.Bind(&distributed.LocalDistributedProvider{})
	// 绑定 配置 服务提供者
	container.Bind(&config.ConfigProvider{})
	// 绑定 日志 服务提供者
	container.Bind(&log.LogProvider{})
	// 绑定 全链路支持 服务提供者
	container.Bind(&id.IDProvider{})
	container.Bind(&trace.TraceProvider{})

	// 将 HTTP 引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(container); err == nil {
		// 绑定 Kernel 服务提供者
		container.Bind(&kernel.KernelProvider{HttpEngine: engine})
	}
	// 运行root命令
	console.RunRootCommand(container, false)
}
