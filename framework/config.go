package framework

import (
	"github.com/spf13/viper"
)

// ServiceConfig 定义一个服务提供者提供的服务配置项
type ServiceConfig interface {

	// 获取配置对应的viper实例
	Viper() viper.Viper

	ConfigStruct() any
	
}
