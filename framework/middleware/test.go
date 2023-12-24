package middleware

import (
	"fmt"
	"gob/framework"
)

func Test1() framework.ControllerHandler {
	// 使用函数回调
	return func(ctx *framework.Context) error {
		fmt.Println("middleware pre test1")
		ctx.Next() // 调用Next往下调用，会自增contxt.index
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() framework.ControllerHandler {
	// 使用函数回调
	return func(ctx *framework.Context) error {
		fmt.Println("middleware pre test2")
		ctx.Next() // 调用Next往下调用，会自增contxt.index
		fmt.Println("middleware post test2")
		return nil
	}
}

func Test3() framework.ControllerHandler {
	// 使用函数回调
	return func(ctx *framework.Context) error {
		fmt.Println("middleware pre test3")
		ctx.Next()
		fmt.Println("middleware post test3")
		return nil
	}
}
