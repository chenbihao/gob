package id

// ServiceProvider 实现文件 provider.go

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

// IDProvider 服务提供者具体实现方法
type IDProvider struct{}

var _ framework.ServiceProvider = (*IDProvider)(nil)

// Register 注册方法
func (provider *IDProvider) Register(container framework.Container) framework.NewInstance {
	return NewIDService
}

// Boot 启动调用
func (provider *IDProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (provider *IDProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (provider *IDProvider) Params(container framework.Container) []interface{} {
	return []interface{}{}
}

// Name 获取字符串凭证
func (provider *IDProvider) Name() string {
	return contract.IDKey
}
