package demo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/gin"
	"github.com/chenbihao/gob/framework/provider/redis"
	"time"
)

// DemoRedis redis的路由方法
func (api *DemoApi) DemoRedis(c *gin.Context) {
	logger := c.MustMakeLog()
	logger.Info(c, "request start", nil)

	// 初始化一个redis
	redisService := c.MustMake(contract.RedisKey).(contract.RedisService)
	client, err := redisService.GetClient(redis.WithConfigPath("cache.default"), redis.WithRedisConfig(func(options *contract.RedisConfig) {
		options.MaxRetries = 3
	}))
	if err != nil {
		logger.Error(c, err.Error(), nil)
		c.AbortWithError(50001, err)
		return
	}
	if err := client.Set(c, "foo", "bar", 1*time.Hour).Err(); err != nil {
		c.AbortWithError(500, err)
		return
	}
	val := client.Get(c, "foo").String()
	logger.Info(c, "redis get", map[string]interface{}{
		"val": val,
	})

	if err := client.Del(c, "foo").Err(); err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, "ok")
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, p)
	if err != nil {
		return fmt.Errorf("解码失败: %w", err)
	}
	return nil
}
func (p Person) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("编码失败: %w", err)
	}
	return data, nil
}

// DemoCache cache的简单例子
func (api *DemoApi) DemoCache(c *gin.Context) {
	logger := c.MustMakeLog()
	logger.Info(c, "request start", nil)
	// 初始化cache服务
	cacheService := c.MustMake(contract.CacheKey).(contract.CacheService)
	// 设置key为foo
	err := cacheService.Set(c, "foo", "bar", 1*time.Hour)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	// 获取key为foo
	val, err := cacheService.Get(c, "foo")
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	logger.Info(c, "cache get", map[string]interface{}{
		"val": val,
	})
	// 删除key为foo
	if err := cacheService.Del(c, "foo"); err != nil {
		c.AbortWithError(500, err)
		return
	}

	// 定义回调函数，用于生成缓存内容
	rememberFunc := func(ctx context.Context, container framework.Container) (interface{}, error) {
		// 在这里可以查询数据库或执行其他操作来获取缓存的内容
		// 这里仅做示例，直接将内容赋值给模型对象
		model := Person{}
		model.Name = "Alice"
		model.Age = 25
		return model, nil
	}

	// 创建一个模型对象，用于存储缓存的内容
	model := &Person{}
	err = cacheService.Remember(c, "fooR", 1*time.Hour, rememberFunc, model)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	logger.Info(c, "cache remember", map[string]interface{}{
		"model": fmt.Sprintf("%+v", model),
	})
	c.JSON(200, model)
}
