package contract

import (
	"fmt"
	"github.com/chenbihao/gob/framework"
	"github.com/redis/go-redis/v9"
)

/*
## 服务介绍：
提供 redis 服务的服务，可以用于获取 redis 连接实例。
## 支持命令：无
## 支持配置：

使用之前需要确保已经正确配置了redis服务。

配置文件为 `config/[env]/redis.yaml`，以下是一个配置的例子：

```yaml
timeout: 10s # 连接超时
read_timeout: 2s # 读超时
write_timeout: 2s # 写超时
write:
	host: localhost # ip地址
	port: 6379 # 端口
	db: 0 #db
	username: # 用户名
	password: "5233" # 密码
	timeout: 10s # 连接超时
	read_timeout: 2s # 读超时
	write_timeout: 2s # 写超时
	conn_min_idle: 10 # 连接池最小空闲连接数
	conn_max_open: 20 # 连接池最大连接数
	conn_max_lifetime: 1h # 连接数最大生命周期
	conn_max_idletime: 1h # 连接数空闲时长
```
*/

// RedisKey 定义字符串凭证
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
