package app

// ServiceProvider 实现文件 provider.go

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

// GobAppProvider 提供App的具体实现方法
type GobAppProvider struct {
	BaseFolder string
}

var _ framework.ServiceProvider = (*GobAppProvider)(nil)

// Register 注册 GobApp 方法
func (appProvider *GobAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewGobApp
}

// Boot 启动调用
func (appProvider *GobAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (appProvider *GobAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (appProvider *GobAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, appProvider.BaseFolder}
}

// Name 获取字符串凭证
func (appProvider *GobAppProvider) Name() string {
	return contract.AppKey
}
