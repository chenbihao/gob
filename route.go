package main

import (
	"github.com/chenbihao/gob/framework/gin"
	"github.com/chenbihao/gob/framework/middleware"
	"time"
)

// 注册路由规则
// func registerRouter(core *framework.Core) {
func registerRouter(core *gin.Engine) {

	core.Use(gin.Recovery()) // 使用 gin 的 Recovery 中间件
	core.Use(middleware.Cost())

	// gin.Engine 的方法为全大写
	core.GET("/user/login", UserLoginController)

	subjectApi := core.Group("/subject")
	{
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		subjectInnerApi.GET("/name", SubjectNameController)
	}
	core.GET("/timeout", middleware.Timeout(10*time.Second), TimeoutController)
}
