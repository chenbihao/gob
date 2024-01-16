package gin

import "github.com/chenbihao/gob/framework"

// --- 服务容器：engine 实现 container 的绑定封装

// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回 error
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

// IsBind 关键字凭证是否已经绑定服务提供者
func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

// SetContainer 设置服务容器
func (engine *Engine) SetContainer(container framework.Container) {
	engine.container = container
}

// GetContainer 从Engine中获取container
func (engine *Engine) GetContainer() framework.Container {
	return engine.container
}
