package main

import (
	"gob/framework"
	"gob/framework/middleware"
	"time"
)

// 注册路由规则
func registerRouter(core *framework.Core) {

	// 扩展需求：core中使用use注册全局中间件 （需放在前面）
	core.Use(middleware.Recovery(), middleware.Cost())

	// 需求1+2：HTTP方法+静态路由匹配
	// 扩展需求：在core中使用middleware.Test3() 为单个路由增加中间件
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// 需求3：批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 需求4:动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}

		// 扩展需求：在 group 中使用 middleware.Test3() 为单个路由增加中间件
		subjectApi.Get("/middleware/test3", middleware.Test3(), SubjectAddController)
	}
	core.Get("/timeout", middleware.Timeout(1*time.Second), TimeoutController)
}
