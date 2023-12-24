package main

import (
	"gob/framework"
	"net/http"
)

func main() {

	// 核心框架初始化
	core := framework.NewCore()

	// 设置路由
	registerRouter(core)

	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8080",
	}
	server.ListenAndServe()
}
