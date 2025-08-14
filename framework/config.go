package framework

import "github.com/knadh/koanf/v2"

// ServiceConfig 定义一个服务提供者提供的服务配置项
type ServiceConfig interface {

	// Koanf *Koanf ？
	Koanf() koanf.Koanf

	// ConfigName 配置名（用来做前缀、独立配置文件）
	ConfigName() string

	// ConfigStruct 配置值
	ConfigStruct() any
}
