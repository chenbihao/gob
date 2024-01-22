package contract

import (
	"fmt"
	"github.com/chenbihao/gob/framework"
	"github.com/redis/go-redis/v9"
)

const RedisKey = "gob:redis"

// RedisService 表示一个redis服务
type RedisService interface {
	// GetClient 获取redis连接实例
	GetClient(option ...RedisOption) (*redis.Client, error)
}

// RedisOption 代表初始化的时候的选项
type RedisOption func(container framework.Container, config *RedisConfig) error

// RedisConfig 为gob定义的Redis配置结构
type RedisConfig struct {
	*redis.Options
}

// UniqKey 用来唯一标识一个RedisConfig配置
func (config *RedisConfig) UniqKey() string {
	return fmt.Sprintf("%v_%v_%v_%v", config.Addr, config.DB, config.Username, config.Network)
}
