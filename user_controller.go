package main

import (
	"github.com/chenbihao/gob/framework/gin"
	"time"
)

func UserLoginController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, UserLoginController")
}

func TimeoutController(c *gin.Context) {
	// 执行具体的业务逻辑
	time.Sleep(8 * time.Second)
	c.ISetOkStatus().IJson("ok, TimeoutController")
}
