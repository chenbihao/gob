package http

import (
	"github.com/chenbihao/gob/app/http/module/demo"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/gin"
	"github.com/chenbihao/gob/framework/middleware"
	"github.com/chenbihao/gob/framework/middleware/cors"
	ginSwagger "github.com/chenbihao/gob/framework/middleware/gin-swagger"
	"github.com/chenbihao/gob/framework/middleware/static"
	swaggerFiles "github.com/swaggo/files"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	container := r.GetContainer()
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	// /路径先去./dist目录下查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./gob_frontend/dist", false)))

	// 使用全链路追踪
	r.Use(middleware.Trace())
	// 使用中间件迁移工具迁移下来的 cors 中间件
	r.Use(cors.Default())

	// 如果配置了swagger，则显示swagger的中间件
	if configService.GetBool("app.swagger") == true {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	demo.Register(r) // 这个demo是业务App自定义的demo服务,位置在 `app/http/module/demo/*`
}
