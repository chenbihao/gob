package middleware

import (
	"context"
	"fmt"
	"gob/framework"
	"log"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	// 使用函数回调
	return func(ctx *framework.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		// 执行业务逻辑前预操作：初始化超时context
		durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 使用next执行具体的业务逻辑
			ctx.Next()

			finish <- struct{}{}
		}()
		// 执行业务逻辑后操作
		select {
		case p := <-panicChan:
			ctx.Json(500, "time out")
			log.Println(p)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			ctx.Json(500, "time out")
			ctx.SetHasTimeout()
		}
		return nil
	}
}
