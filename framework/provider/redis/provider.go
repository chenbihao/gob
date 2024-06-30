package redis

// ServiceProvider 实现文件 provider.go

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

type RedisProvider struct {
	c framework.Container
}

var _ framework.ServiceProvider = (*RedisProvider)(nil)

func (provider *RedisProvider) Name() string {
	return contract.RedisKey
}

func (provider *RedisProvider) Register(container framework.Container) framework.NewInstance {
	return NewRedisService
}

func (provider *RedisProvider) IsDefer() bool {
	return true
}

func (provider *RedisProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

func (provider *RedisProvider) Boot(container framework.Container) error {
	return nil
}
