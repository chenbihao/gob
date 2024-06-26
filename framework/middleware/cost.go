package middleware

import (
	"github.com/chenbihao/gob/framework/gin"
	"log"
	"time"
)

// Cost struct
type Cost struct {
}

// Func 计时
func (c *Cost) Func() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		// 记录开始时间
		start := time.Now()

		log.Printf("api uri start: %v", c.Request.RequestURI)
		// 使用next执行具体的业务逻辑
		c.Next()

		// 记录结束时间
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri end: %v, cost: %v", c.Request.RequestURI, cost.Seconds())

	}
}
