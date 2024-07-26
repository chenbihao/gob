package log

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/provider/log/services"
	"strings"
)

// LogProvider 服务提供者具体实现方法
type LogProvider struct {
	// 运行模式
	Drivers []string
	// 日志context上下文信息获取函数
	CtxFielder contract.CtxFielder
}

var _ framework.ServiceProvider = (*LogProvider)(nil)

// Register 注册 GobAppService 方法
func (provider *LogProvider) Register(container framework.Container) framework.NewInstance {

	return services.NewLogService
}

// Boot 启动调用
func (provider *LogProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (provider *LogProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (provider *LogProvider) Params(container framework.Container) []interface{} {
	// 获取configService
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	// 设置参数 Drivers
	if len(provider.Drivers) == 0 {
		for _, driver := range configService.GetStringSlice("log.drivers") {
			provider.Drivers = append(provider.Drivers, strings.ToLower(driver))
		}
	}
	// 定义5个参数
	return []interface{}{container, provider.Drivers, provider.CtxFielder}
}

// Name 获取字符串凭证
func (provider *LogProvider) Name() string {
	return contract.LogKey
}
