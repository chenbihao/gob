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

func (sp *RedisProvider) Name() string {
	return contract.RedisKey
}

func (sp *RedisProvider) Register(c framework.Container) framework.NewInstance {
	return NewRedisService
}

func (sp *RedisProvider) IsDefer() bool {
	return true
}

func (sp *RedisProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *RedisProvider) Boot(c framework.Container) error {
	return nil
}
