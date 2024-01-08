package env

// ServiceProvider 实现文件 provider.go

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

// EnvProvider 提供App的具体实现方法
type EnvProvider struct {
	Folder string
}

var _ framework.ServiceProvider = (*EnvProvider)(nil)

// Register 注册 GobAppService 方法
func (provider *EnvProvider) Register(container framework.Container) framework.NewInstance {
	return NewGobEnvService
}

// Boot 启动调用
func (provider *EnvProvider) Boot(container framework.Container) error {
	app := container.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer 是否延迟初始化
func (provider *EnvProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (provider *EnvProvider) Params(container framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

// Name 获取字符串凭证
func (provider *EnvProvider) Name() string {
	return contract.EnvKey
}
