---
lang: zh-CN
title: gob:env
description:
---
# gob:env

## 服务介绍：
提供环境变量相关方法
## 支持命令：无
## 支持配置：无

## 使用方法
```go 
type Env interface {
	// AppEnv 获取当前的环境，建议分为 dev/test/prod
	AppEnv() string
	// IsExist 判断一个环境变量是否有被设置
	IsExist(string) bool
	// Get 获取某个环境变量，如果没有设置，返回""
	Get(string) string
	// All 获取所有的环境变量，.env 和运行环境变量融合后结果
	All() map[string]string
}
```
