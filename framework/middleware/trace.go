package middleware

import (
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/gin"
)

// Trace ...
type Trace struct{}

// NewTrace ...
func NewTrace() *Trace {
	return &Trace{}
}

// Func 全链路ID
func (t *Trace) Func() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		tracer := c.MustMake(contract.TraceKey).(contract.Trace)
		traceCtx := tracer.ExtractHTTP(c.Request)
		tracer.WithTrace(c, traceCtx)

		// 使用next执行具体的业务逻辑
		c.Next()
	}
}
