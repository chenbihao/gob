---
lang: zh-CN
title: gob:config
description:
---
# gob:config

## 服务介绍：
提供基础的配置文件获取方法
## 支持命令：
[app](../command/config)
## 支持配置：无

## 提供方法：
```go 
type Config interface {
	// IsExist 检查一个属性是否存在
	IsExist(key string) bool
	// GetIntIfExist 如果存在的话，获取一个 int 属性
	GetIntIfExist(key string) (exist bool, value int)
	// GetStringIfExist 如果存在的话，获取一个 string 属性
	GetStringIfExist(key string) (exist bool, value string)
	// Get 获取一个属性值
	Get(key string) interface{}
	// GetBool 获取一个 bool 属性
	GetBool(key string) bool
	// GetInt 获取一个 int 属性
	GetInt(key string) int
	// GetFloat64 获取一个 float64 属性
	GetFloat64(key string) float64
	// GetTime 获取一个 time 属性
	GetTime(key string) time.Time
	// GetString 获取一个 string 属性
	GetString(key string) string
	// GetIntSlice 获取一个 int 数组属性
	GetIntSlice(key string) []int
	// GetStringSlice 获取一个 string 数组
	GetStringSlice(key string) []string
	// GetStringMap 获取一个 string 为 key，interface 为 val 的 map
	GetStringMap(key string) map[string]interface{}
	// GetStringMapString 获取一个 string 为 key，string 为 val 的 map
	GetStringMapString(key string) map[string]string
	// GetStringMapStringSlice 获取一个 string 为 key，数组 string 为 val 的 map
	GetStringMapStringSlice(key string) map[string][]string
	// Load 加载配置到某个对象
	Load(key string, val interface{}) error
}
```
