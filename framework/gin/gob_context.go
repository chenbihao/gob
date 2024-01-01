package gin

import (
	"context"
	"github.com/chenbihao/gob/framework"
)

// --- 基础能力

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

// --- 服务容器：engine 实现 container 的绑定封装

// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回 error
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

// IsBind 关键字凭证是否已经绑定服务提供者
func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

// --- 服务容器：context 实现 container 的几个封装

// 实现 make 的封装
func (ctx *Context) Make(key string) (interface{}, error) {
	return ctx.container.Make(key)
}

// 实现 mustMake 的封装
func (ctx *Context) MustMake(key string) interface{} {
	return ctx.container.MustMake(key)
}

// 实现 makeNew 的封装
func (ctx *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return ctx.container.MakeNew(key, params)
}
