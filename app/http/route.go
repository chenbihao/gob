package http

import (
	"github.com/chenbihao/gob/app/http/module/demo"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/gin"
	"github.com/chenbihao/gob/framework/middleware"
	"github.com/chenbihao/gob/framework/middleware/gzip"
	"github.com/chenbihao/gob/framework/middleware/static"

	// docs "github.com/go-project-name/docs"
	// 项目生成的docs，这里是 _ "github.com/chenbihao/gob/app/http/swagger"
	// 放在  `app/http/swagger.go` 那边

	ginSwagger "github.com/chenbihao/gob/framework/middleware/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	container := r.GetContainer()
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	// `/` 路径先去 `./dist` 目录下查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./gob_frontend/dist", false)))

	// 使用全链路追踪
	r.Use(middleware.NewTrace().Func())
	// 使用中间件迁移工具迁移下来的 cors 中间件
	r.Use(middleware.NewCORS().Func())
	// 使用中间件迁移工具迁移下来的 gzip 中间件
	r.Use(gzip.Gzip(gzip.BestSpeed))

	// 如果配置了swagger，则显示swagger的中间件
	if configService.GetBool("app.swagger") {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	demo.Register(r) // 这个demo是业务App自定义的demo服务,位置在 `app/http/module/demo/*`
}
