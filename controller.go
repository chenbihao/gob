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
		time.Sleep(2 * time.Second)

		c.Json(200, "ok") // todo 这里如果同时超时，没上锁也会导致问题
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
		c.SetHasTimeout()
	}
	return nil
}

// 以下是封装context的作用

// func Foo(request *http.Request, response http.ResponseWriter) {
// 	obj := map[string]interface{}{
// 		"errno":  50001,
// 		"errmsg": "inner error",
// 		"data":   nil,
// 	}

// 	response.Header().Set("Content-Type", "application/json")

// 	foo := request.PostFormValue("foo")
// 	if foo == "" {
// 		foo = "10"
// 	}
// 	fooInt, err := strconv.Atoi(foo)
// 	if err != nil {
// 		response.WriteHeader(500)
// 		return
// 	}
// 	obj["data"] = fooInt
// 	byt, err := json.Marshal(obj)
// 	if err != nil {
// 		response.WriteHeader(500)
// 		return
// 	}
// 	response.WriteHeader(200)
// 	response.Write(byt)
// 	return
// }

// func Foo2(ctx *framework.Context) error {
// 	obj := map[string]interface{}{
// 		"errno":  50001,
// 		"errmsg": "inner error",
// 		"data":   nil,
// 	}

// 	fooInt := ctx.FormInt("foo", 10)
// 	obj["data"] = fooInt
// 	return ctx.Json(http.StatusOK, obj)
// }

// func Foo3(ctx *framework.Context) error {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})

// 	return rdb.Set(ctx, "key", "value", 0).Err()
// }
