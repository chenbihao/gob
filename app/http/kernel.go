package http

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/gin"
)

// NewHttpEngine 创建了一个绑定了路由的 Web 引擎
func NewHttpEngine(container framework.Container) (*gin.Engine, error) {

	// 设置为 Release，为的是默认在启动中不输出调试信息
	// if container.MustMake(contract.EnvKey).(contract.Env).AppEnv() == contract.EnvProd {}
	gin.SetMode(gin.ReleaseMode)

	// 默认启动一个 Web 引擎 （Default 包含 Logger and Recovery）
	r := gin.Default()

	// 设置了 Engine
	r.SetContainer(container)

	// 业务绑定路由操作
	Routes(r)

	// 返回绑定路由后的 Web 引擎
	return r, nil
}
