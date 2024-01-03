package gin

import (
	"context"
)

// --- 基础能力

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
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
