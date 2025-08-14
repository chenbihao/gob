package config

// ServiceProvider 实现文件 provider.go

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

// ConfigProvider 服务提供者具体实现方法
type ConfigProvider struct {
}

var _ framework.ServiceProvider = (*ConfigProvider)(nil)

// Name 定义对应的服务字符串凭证
func (provider *ConfigProvider) Name() string {
	return contract.ConfigKey
}

// IsDefer 是否延迟加载
func (provider *ConfigProvider) IsDefer() bool {
	return false
}

// Boot 启动的时候注入
func (provider *ConfigProvider) Boot(container framework.Container) error {
	return nil
}

// Params 定义要传递给实例化方法的参数
func (provider *ConfigProvider) Params(container framework.Container) []interface{} {
	return []any{container}
}

// Register 注册一个服务实例
func (provider *ConfigProvider) Register(container framework.Container) framework.NewInstance {
	return NewConfigService
}
