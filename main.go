package main

import (
	"github.com/chenbihao/gob/app/http"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/command"
	"github.com/chenbihao/gob/framework/provider/app"
	"github.com/chenbihao/gob/framework/provider/cache"
	"github.com/chenbihao/gob/framework/provider/config"
	"github.com/chenbihao/gob/framework/provider/distributed"
	"github.com/chenbihao/gob/framework/provider/env"
	"github.com/chenbihao/gob/framework/provider/id"
	"github.com/chenbihao/gob/framework/provider/kernel"
	"github.com/chenbihao/gob/framework/provider/log"
	"github.com/chenbihao/gob/framework/provider/orm"
	"github.com/chenbihao/gob/framework/provider/redis"
	"github.com/chenbihao/gob/framework/provider/sls"
	"github.com/chenbihao/gob/framework/provider/ssh"
	"github.com/chenbihao/gob/framework/provider/trace"
)

func main() {

	// 初始化服务容器
	container := framework.NewGobContainer()
	_ = container.Bind(&app.AppProvider{})                      // 绑定 App 服务提供者
	_ = container.Bind(&env.EnvProvider{})                      // 绑定 环境变量 服务提供者
	_ = container.Bind(&distributed.LocalDistributedProvider{}) // 绑定 本地式分布锁 服务提供者
	_ = container.Bind(&config.ConfigProvider{})                // 绑定 配置 服务提供者
	_ = container.Bind(&log.LogProvider{})                      // 绑定 日志 服务提供者
	_ = container.Bind(&id.IDProvider{})                        // 绑定 ID 服务提供者
	_ = container.Bind(&trace.TraceProvider{})                  // 绑定 trace 服务提供者
	_ = container.Bind(&orm.GormProvider{})                     // 绑定 orm 服务提供者
	_ = container.Bind(&redis.RedisProvider{})                  // 绑定 redis 服务提供者
	_ = container.Bind(&cache.CacheProvider{})                  // 绑定 缓存 服务提供者
	_ = container.Bind(&ssh.SSHProvider{})                      // 绑定 ssh 服务提供者
	_ = container.Bind(&sls.SLSProvider{})                      // 绑定 sls 服务提供者

	// 将 HTTP 引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(container); err == nil {
		// 绑定 Kernel 服务提供者
		_ = container.Bind(&kernel.KernelProvider{HttpEngine: engine})
	}
	// 运行root命令
	_ = command.RunRootCommand(container)
}
