package app

// ServiceProvider 实现文件 provider.go

import (
	"fmt"
	"log"

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
	// 注册 AppConfig 到 ConfigService
	// 使用 Make 而不是 MustMake，避免在 Config 服务未注册时 panic
	configService, err := container.Make(contract.ConfigKey)
	if err != nil {
		// Config 服务不可用，使用默认值初始化 App
		return nil
	}

	// 使用 SubConfigRegistry 接口注册配置，避免循环依赖和封装违反
	if subConfigRegistry, ok := configService.(contract.SubConfigRegistry); ok {
		if err := subConfigRegistry.RegisterSubConfigBySubConfigRegistry(&AppConfig{}); err != nil {
			return fmt.Errorf("failed to register app config: %w", err)
		}
	} else {
		// 如果不支持子配置注册，记录警告但不阻止启动
		log.Printf("warning: config service does not support sub-config registration")
	}
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
