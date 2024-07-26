package demo

import (
	demoService "github.com/chenbihao/gob/app/provider/demo"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/gin"
)

type DemoApi struct {
	service *Service
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	_ = r.Bind(&demoService.DemoProvider{})

	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.POST("/demo/demo_post", api.DemoPost)

	r.GET("/demo/orm", api.DemoOrm)
	r.GET("/demo/redis", api.DemoRedis)
	r.GET("/demo/cache", api.DemoCache)

	r.GET("/demo/logs", api.DemoLogs)
	return nil
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{service: service}
}

// Demo godoc
// @Summary 获取所有用户
// @Description 获取所有用户
// @Produce  json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /demo/demo [get]
func (api *DemoApi) Demo(c *gin.Context) {
	// 获取password
	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	password := configService.GetString("database.mysql.password")

	logger := c.MustMakeLog()
	logger.Info(c, "demo test logger", map[string]interface{}{
		"api":      "demo/demo",
		"password": password,
	})

	// 打印出来
	c.JSON(200, "测试dev热更新模式 123444")
}

// Demo godoc
// @Summary 获取所有学生
// @Description 获取所有学生
// @Produce  json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /demo/demo2 [get]
func (api *DemoApi) Demo2(c *gin.Context) {
	demoProvider := c.MustMake(demoService.DemoKey).(demoService.IService)
	students := demoProvider.GetAllStudent()
	usersDTO := StudentsToUserDTOs(students)
	c.JSON(200, usersDTO)
}

func (api *DemoApi) DemoPost(c *gin.Context) {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := c.BindJSON(&foo)
	if err != nil {
		_ = c.AbortWithError(500, err)
	}
	c.JSON(200, nil)
}

func (api *DemoApi) DemoLogs(c *gin.Context) {
	log := c.MustMakeLog()
	count, _ := c.DefaultQueryInt("count", 1)
	for i := 0; i < count; i++ {
		log.Error(c, "Error( ErrorErrorError )", nil)
		log.Warn(c, "Warn( WarnWarnWarn )", nil)
		log.Info(c, "Info( InfoInfoInfo )", nil)
		log.Info(c, "Info( InfoInfoInfo )", map[string]interface{}{"info": "xxxxxxxxxxxxx"})
		log.Debug(c, "Debug( DebugDebugDebug )", nil)
		log.Trace(c, "Trace( TraceTraceTrace )", nil)
	}
	c.JSON(200, "已打印测试日志")
}
