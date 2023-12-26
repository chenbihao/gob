package main

import (
	"gob/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, UserLoginController")
	return nil
}

func TimeoutController(c *framework.Context) error {
	// 执行具体的业务逻辑
	time.Sleep(2 * time.Second)
	c.SetOkStatus().Json("ok, TimeoutController")
	return nil
}
