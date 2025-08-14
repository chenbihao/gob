package config_old

// ServiceProvider 实现文件 provider.go

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

// OldConfigProvider 服务提供者具体实现方法
type OldConfigProvider struct {
}

var _ framework.ServiceProvider = (*OldConfigProvider)(nil)

// Name 定义对应的服务字符串凭证
func (provider *OldConfigProvider) Name() string {
	return contract.OldConfigKey
}

// IsDefer 是否延迟加载
func (provider *OldConfigProvider) IsDefer() bool {
	return false
}

// Boot 启动的时候注入
func (provider *OldConfigProvider) Boot(container framework.Container) error {
	return nil
}

// Params 定义要传递给实例化方法的参数
func (provider *OldConfigProvider) Params(container framework.Container) []interface{} {
	appService := container.MustMake(contract.AppKey).(contract.App)
	//envService := container.MustMake(contract.EnvKey).(contract.Env)
	//appEnv := envService.AppEnv()
	//configFolder := appService.ConfigFolder()
	//appEnvFolder := filepath.Join(configFolder, appEnv)
	return []any{container, appService.ConfigFolder()}
}

// Register 注册一个服务实例
func (provider *OldConfigProvider) Register(container framework.Container) framework.NewInstance {
	return NewConfigService
}
