---
lang: zh-CN
title: gob:trace
description:
---
# gob:trace

## 服务介绍：
提供分布式链路追踪服务，可以用于跟踪分布式服务调用链路。
## 支持命令：无
## 支持配置：无

## 提供方法：
```go 
type Trace interface {
	// WithTrace register new trace to context
	WithTrace(c context.Context, trace *TraceContext) context.Context
	// GetTrace From trace context
	GetTrace(c context.Context) *TraceContext
	// NewTrace generate a new trace
	NewTrace() *TraceContext
	// StartSpan generate cspan for child call
	StartSpan(trace *TraceContext) *TraceContext

	// ToMap traceContext to map for logger
	ToMap(trace *TraceContext) map[string]string

	// ExtractHTTP GetTrace By Http
	ExtractHTTP(req *http.Request) *TraceContext
	// InjectHTTP Set Trace to Http
	InjectHTTP(req *http.Request, trace *TraceContext) *http.Request
}
```
