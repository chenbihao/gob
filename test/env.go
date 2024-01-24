package test

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/provider/app"
	"github.com/chenbihao/gob/framework/provider/config"
	"github.com/chenbihao/gob/framework/provider/env"
	"github.com/chenbihao/gob/framework/provider/log"
)

const (
	BasePath = "D:\\DevProjects\\自己库\\gob\\"
)

func InitBaseContainer() framework.Container {
	// 初始化服务容器
	container := framework.NewGobContainer()
	// 绑定App服务提供者
	container.Bind(&app.AppProvider{BaseFolder: BasePath})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.GobTestingEnvProvider{})
	return container
}

func InitBaseConfLogContainer() framework.Container {
	container := InitBaseContainer()
	container.Bind(&config.ConfigProvider{})
	container.Bind(&log.LogProvider{})
	return container
}
