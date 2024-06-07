---
lang: zh-CN
title: gob:id
description:
---
# gob:id

## 服务介绍：
提供分布式 ID 生成服务，可以为当前服务生成唯一id
## 支持命令：无
## 支持配置：无

## 提供方法：
```go 
type ID interface {
	NewID() string
}
```
