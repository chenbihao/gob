package main

import (
	"fmt"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/gin"
	"github.com/chenbihao/gob/framework/provider/app"
)

// 对应路由 /subject/list/all
func SubjectListController(c *gin.Context) {
	// 获取 App 服务实例
	appService := c.MustMake(contract.AppKey).(app.GobApp)

	// 输出结果
	c.ISetOkStatus().IJson(appService.ConfigFolder())
}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectUpdateController")
}

func SubjectGetController(c *gin.Context) {
	subjectId, _ := c.DefaultParamInt("id", 0)
	c.ISetOkStatus().IJson("ok, SubjectGetController:" + fmt.Sprint(subjectId))
}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")
}
