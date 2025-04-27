package app

// ServiceProvider 实现文件 provider.go

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

// AppProvider 服务提供者具体实现方法
type AppProvider struct {
	BaseFolder string
}

var _ framework.ServiceProvider = (*AppProvider)(nil)

// Name 获取字符串凭证
func (provider *AppProvider) Name() string {
	return contract.AppKey
}

// IsDefer 是否延迟初始化
func (provider *AppProvider) IsDefer() bool {
	return false
}

// Boot 启动调用
func (provider *AppProvider) Boot(container framework.Container) error {
	return nil
}

// Params 获取初始化参数
func (provider *AppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, provider.BaseFolder}
}

// Register 注册 AppService 方法
func (provider *AppProvider) Register(container framework.Container) framework.NewInstance {
	return NewGobApp
}
