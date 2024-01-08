package main

import (
	"github.com/chenbihao/gob/app/console"
	"github.com/chenbihao/gob/app/http"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/provider/app"
	"github.com/chenbihao/gob/framework/provider/distributed"
	"github.com/chenbihao/gob/framework/provider/env"
	"github.com/chenbihao/gob/framework/provider/kernel"
)

func main() {

	// 初始化服务容器
	container := framework.NewGobContainer()
	// 绑定 App 服务提供者
	container.Bind(&app.GobAppProvider{})
	// 绑定 环境变量 服务提供者
	container.Bind(&env.EnvProvider{})
	// 绑定 分布式本地锁 服务提供者
	container.Bind(&distributed.LocalDistributedProvider{})

	// 将 HTTP 引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		// 绑定 Kernel 服务提供者
		container.Bind(&kernel.GobKernelProvider{HttpEngine: engine})
	}
	// 运行root命令
	console.RunCommand(container)
}
