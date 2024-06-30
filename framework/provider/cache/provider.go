package cache

// ServiceProvider 实现文件 provider.go

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/provider/cache/services"
	"strings"
)

// CacheProvider 服务提供者
type CacheProvider struct {
	Driver string // Driver
}

var _ framework.ServiceProvider = (*CacheProvider)(nil)

// Register 注册一个服务实例
func (provider *CacheProvider) Register(container framework.Container) framework.NewInstance {
	if provider.Driver == "" {
		tcs, err := container.Make(contract.ConfigKey)
		if err != nil {
			// 默认使用 内存模式
			return services.NewMemoryCache
		}

		cs := tcs.(contract.Config)
		provider.Driver = strings.ToLower(cs.GetString("cache.driver"))
	}

	// 根据driver的配置项确定
	switch provider.Driver {
	case "redis":
		return services.NewRedisCache
	case "memory":
		return services.NewMemoryCache
	default:
		return services.NewMemoryCache
	}
}

// Boot 启动的时候注入
func (provider *CacheProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟加载
func (provider *CacheProvider) IsDefer() bool {
	return true
}

// Params 定义要传递给实例化方法的参数
func (provider *CacheProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

// Name 定义对应的服务字符串凭证
func (provider *CacheProvider) Name() string {
	return contract.CacheKey
}
