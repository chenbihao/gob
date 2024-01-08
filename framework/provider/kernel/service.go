package kernel

import (
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/gin"
	"net/http"
)

// 引擎服务
type KernelService struct {
	engine *gin.Engine
}

var _ contract.Kernel = (*KernelService)(nil)

// 初始化 web 引擎服务实例
func NewGobKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &KernelService{engine: httpEngine}, nil
}

// 返回 web 引擎
func (s *KernelService) HttpEngine() http.Handler {
	return s.engine
}
