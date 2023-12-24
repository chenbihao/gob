package main

import (
	"gob/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	c.Json(200, "ok, UserLoginController")
	return nil
}

func TimeoutController(c *framework.Context) error {
	// 执行具体的业务逻辑
	time.Sleep(2 * time.Second)
	c.Json(200, "ok, TimeoutController")
	return nil
}
