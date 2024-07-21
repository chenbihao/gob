package log

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/provider/log/formatter"
	"github.com/chenbihao/gob/framework/provider/log/services"
	"io"
	"strings"
)

// LogProvider 服务提供者具体实现方法
type LogProvider struct {
	// 运行模式
	Driver string
	// 日志级别
	Level contract.LogLevel
	// 日志输出格式方法
	Formatter contract.Formatter
	// 日志context上下文信息获取函数
	CtxFielder contract.CtxFielder
	// 日志输出信息
	Output io.Writer
}

var _ framework.ServiceProvider = (*LogProvider)(nil)

// Register 注册 GobAppService 方法
func (provider *LogProvider) Register(container framework.Container) framework.NewInstance {





	if provider.Driver == "" {
		configContainer := container.MustMake(contract.ConfigKey)
		config := configContainer.(contract.Config)
		provider.Driver = strings.ToLower(config.GetString("log.driver"))
	}
	// 根据driver的配置项确定
	switch provider.Driver {
	case "single":
		return services.NewSingleLogService
	case "rotate":
		return services.NewRotateLogService
	case "console":
		return services.NewConsoleLogService
	case "aliyun_sls":
		return services.NewSlsLog
	default:
		return services.NewConsoleLogService
	}
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
	// 设置参数formatter
	if provider.Formatter == nil {
		provider.Formatter = formatter.TextFormatter
		if configService.IsExist("log.formatter") {
			v := configService.GetString("log.formatter")
			if v == "json" {
				provider.Formatter = formatter.JsonFormatter
			} else if v == "text" {
				provider.Formatter = formatter.TextFormatter
			}
		}
	}
	// 设置参数level
	if provider.Level == contract.UnknownLevel {
		provider.Level = contract.InfoLevel
		if configService.IsExist("log.level") {
			provider.Level = logLevel(configService.GetString("log.level"))
		}
	}
	// 定义5个参数
	return []interface{}{container, provider.Level, provider.CtxFielder, provider.Formatter, provider.Output}
}

// Name 获取字符串凭证
func (provider *LogProvider) Name() string {
	return contract.LogKey
}

// logLevel get level from string
func logLevel(config string) contract.LogLevel {
	switch strings.ToLower(config) {
	case "panic":
		return contract.PanicLevel
	case "fatal":
		return contract.FatalLevel
	case "error":
		return contract.ErrorLevel
	case "warn":
		return contract.WarnLevel
	case "info":
		return contract.InfoLevel
	case "debug":
		return contract.DebugLevel
	case "trace":
		return contract.TraceLevel
	}
	return contract.UnknownLevel
}
