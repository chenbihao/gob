package config

// ServiceProvider 实现文件 provider.go

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"path/filepath"
)

// ConfigProvider 服务提供者具体实现方法
type ConfigProvider struct {
}

var _ framework.ServiceProvider = (*ConfigProvider)(nil)

// Register 注册 GobAppService 方法
func (provider *ConfigProvider) Register(container framework.Container) framework.NewInstance {
	return NewGobConfigService
}

// Boot 启动调用
func (provider *ConfigProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (provider *ConfigProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (provider *ConfigProvider) Params(container framework.Container) []interface{} {
	appService := container.MustMake(contract.AppKey).(contract.App)
	envService := container.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()
	// 配置文件夹地址
	configFolder := appService.ConfigFolder()
	envFolder := filepath.Join(configFolder, env)
	return []interface{}{container, envFolder, envService.All()}
}

// Name 获取字符串凭证
func (provider *ConfigProvider) Name() string {
	return contract.ConfigKey
}
