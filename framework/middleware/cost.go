package middleware

import (
	"gob/framework"
	"log"
	"time"
)

// 计时
func Cost() framework.ControllerHandler {
	// 使用函数回调
	return func(ctx *framework.Context) error {
		// 记录开始时间
		start := time.Now()

		// 使用next执行具体的业务逻辑
		ctx.Next()

		// 记录结束时间
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %v, cost: %.5f", ctx.GetRequest().RequestURI, cost.Seconds())

		return nil
	}
}
