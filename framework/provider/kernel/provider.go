package kernel

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/gin"
)

// GobKernelProvider 提供web引擎
type GobKernelProvider struct {
	HttpEngine *gin.Engine
}

// Register 注册服务提供者
func (provider *GobKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewGobKernelService
}

// Boot 启动的时候判断是否由外界注入了Engine，如果注入的化，用注入的，如果没有，重新实例化
func (provider *GobKernelProvider) Boot(c framework.Container) error {
	if provider.HttpEngine == nil {
		provider.HttpEngine = gin.Default()
	}
	provider.HttpEngine.SetContainer(c)
	return nil
}

// IsDefer 引擎的初始化我们希望开始就进行初始化
func (provider *GobKernelProvider) IsDefer() bool {
	return false
}

// Params 参数就是一个HttpEngine
func (provider *GobKernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.HttpEngine}
}

// Name 提供凭证
func (provider *GobKernelProvider) Name() string {
	return contract.KernelKey
}
