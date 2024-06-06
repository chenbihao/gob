---
lang: zh-CN
title: gob:kernel
description:
---
# gob:kernel

## 服务介绍：
提供框架最核心的结构，包括 http 和 grpc 的 Engine 结构。
## 支持命令：无
## 支持配置：无

## 使用方法
```go 
type Kernel interface {
	// HttpEngine http.Handler结构，作为net/http框架使用, 实际上是gin.Engine
	HttpEngine() http.Handler
}
```
