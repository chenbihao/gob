package sls

import (
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/chenbihao/gob/framework/provider/config"
	tests "github.com/chenbihao/gob/test"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
	"time"
)

func TestSLSService_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&config.ConfigProvider{})

	Convey("test SLSInstance", t, func() {
		slsService, err := NewSLSService(container)
		So(err, ShouldBeNil)
		So(slsService, ShouldNotBeNil)
		service, ok := slsService.(*SLSService)
		So(ok, ShouldBeTrue)
		producerInstance, err := service.GetSLSInstance()
		So(err, ShouldBeNil)
		project, err := service.GetProject()
		So(err, ShouldBeNil)
		So(project, ShouldNotBeNil)
		logstore, err := service.GetLogstore()
		So(err, ShouldBeNil)
		So(logstore, ShouldNotBeNil)

		for i := 0; i < 100; i++ {
			logger := producer.GenerateLog(uint32(time.Now().Unix()),
				map[string]string{"content": "test SLSInstance send balabalabala " + strconv.Itoa(i)})
			err = producerInstance.SendLog(project, logstore, "topic", "127.0.0.1", logger)
			So(err, ShouldBeNil)
		}

		service.producerInstance.SafeClose()
	})
}
