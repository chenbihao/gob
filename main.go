package main

import (
	"gob/framework"
	"gob/framework/middleware"
	"net/http"
)

func main() {

	// 设置路由
	core := framework.NewCore()
	// core中使用use注册中间件
	core.Use(
		middleware.Test1(),
		middleware.Test2())

	// group中使用use注册中间件
	subjectApi := core.Group("/subject")
	subjectApi.Use(middleware.Test3())

	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8080",
	}
	server.ListenAndServe()
}
