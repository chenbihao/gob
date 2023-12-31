package http

import (
	"github.com/chenbihao/gob/app/http/module/demo"
	"github.com/chenbihao/gob/framework/gin"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r) // 这个demo是业务App自定义的demo服务,位置在 `app/http/module/demo/*`
}
