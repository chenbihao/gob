package app

import (
	"fmt"

	"github.com/chenbihao/gob/framework"
)

// AppConfig 实现 ServiceConfig 接口
type AppConfig struct {
	Debug   bool   `koanf:"debug"`
	Version string `koanf:"version"`
}

var _ framework.ServiceConfig = (*AppConfig)(nil)

// ConfigName 返回配置名称，对应配置文件中的 key
func (a *AppConfig) ConfigName() string {
	return "app"
}

// ConfigStruct 返回配置结构体实例，用于反序列化
func (a *AppConfig) ConfigStruct() interface{} {
	return &AppConfig{}
}

// Defaults 返回默认配置值
func (a *AppConfig) Defaults() map[string]interface{} {
	return map[string]interface{}{
		"debug":   false,
		"version": "1.0.0",
	}
}

// Validate 验证配置是否有效
func (a *AppConfig) Validate(config interface{}) error {
	cfg, ok := config.(*AppConfig)
	if !ok {
		return fmt.Errorf("invalid config type for AppConfig")
	}
	if cfg.Version == "" {
		return fmt.Errorf("version is required")
	}
	return nil
}
