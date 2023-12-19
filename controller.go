package main

import (
	"context"
	"fmt"
	"gob/framework"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {

	// 生成一个超时的 Context
	durationCtx, cancel := context.WithTimeout(c.BaseContext(), 1*time.Second)
	// 当所有事情处理结束后调用 cancel，告知 durationCtx 的后续 Context 结束
	defer cancel()

	finish := make(chan struct{}, 1)       // 这个 channel 负责通知结束
	panicChan := make(chan interface{}, 1) // 这个 channel 负责通知 panic 异常

	// 创建一个新的 Goroutine 来处理业务逻辑
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		// 这里做具体的业务
		time.Sleep(1 * time.Second)

		c.Json(200, "ok")
		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimeout() // 这里需要考虑到与业务输出没写保护的问题
	}
	return nil
}
