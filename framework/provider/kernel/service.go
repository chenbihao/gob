package kernel

import (
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/gin"
	"net/http"
)

// 引擎服务
type GobKernelService struct {
	engine *gin.Engine
}

var _ contract.Kernel = (*GobKernelService)(nil)

// 初始化 web 引擎服务实例
func NewGobKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &GobKernelService{engine: httpEngine}, nil
}

// 返回 web 引擎
func (s *GobKernelService) HttpEngine() http.Handler {
	return s.engine
}
