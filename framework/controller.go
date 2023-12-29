package framework

import "github.com/chenbihao/gob/framework/gin"

type ControllerHandler func(c *gin.Context) error
