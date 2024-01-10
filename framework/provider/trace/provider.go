package trace

// ServiceProvider 实现文件 provider.go

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

// TraceProvider 服务提供者具体实现方法
type TraceProvider struct {
	c framework.Container
}

var _ framework.ServiceProvider = (*TraceProvider)(nil)

// Register 注册方法
func (provider *TraceProvider) Register(container framework.Container) framework.NewInstance {
	return NewTraceService
}

// Boot 启动调用
func (provider *TraceProvider) Boot(container framework.Container) error {
	provider.c = container
	return nil
}

// IsDefer 是否延迟初始化
func (provider *TraceProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (provider *TraceProvider) Params(container framework.Container) []interface{} {
	return []interface{}{provider.c}
}

// Name 获取字符串凭证
func (provider *TraceProvider) Name() string {
	return contract.TraceKey
}
