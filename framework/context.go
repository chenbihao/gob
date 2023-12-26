package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

// 自定义 Context
type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter

	hasTimeout bool        // 是否超时标记位
	writerMux  *sync.Mutex // 写保护机制

	handlers []ControllerHandler // 当前请求的handler链条
	index    int                 // 当前请求调用到调用链的哪个节点

	params map[string]string // url路由匹配的参数
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		writerMux:      &sync.Mutex{},
		index:          -1,
	}
}

// base function （封装基本的函数功能，比如获取 http.Request 结构）

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// 为context设置handlers
func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

// 设置参数
func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}

// 核心函数，调用context的下一个函数
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

// 直接返回原生 Context
func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

// implement context.Context （实现标准 Context 接口）

var _ context.Context = new(Context) // 确保某类型实现某接口

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key any) any {
	return ctx.BaseContext().Value(key)
}

// query url
