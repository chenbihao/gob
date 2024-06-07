---
lang: zh-CN
title: gob:app
description:
---
# gob:app

## 服务介绍：
提供基础的 app 框架目录结构获取功能
## 支持命令：无
## 支持配置：无

## 提供方法：
```go 
type App interface {
	// AppID 表示当前这个app的唯一id, 可以用于分布式锁等
	AppID() string
	// Version 定义当前版本
	Version() string
	// BaseFolder 定义项目基础地址
	BaseFolder() string

	// ---------------- 根目录下

	// AppFolder 定义业务代码所在的目录，用于监控文件变更使用
	AppFolder() string
	// ConfigFolder 定义了配置文件的路径
	ConfigFolder() string
	// TestFolder 存放测试所需要的信息
	TestFolder() string
	// StorageFolder 存储文件地址
	StorageFolder() string
	// DeployFolder 存放部署的时候创建的文件夹
	DeployFolder() string

	// ---------------- app 目录下

	// ConsoleFolderr 定义业务自己的命令行服务提供者地址
	ConsoleFolder() string
	// HttpFolderr 定义业务自己的web服务提供者地址
	HttpFolder() string
	// ProviderFolder 定义业务自己的通用服务提供者地址
	ProviderFolder() string

	// ---------------- config 目录下

	// ---------------- storage 目录下

	// LogFolder 定义了日志所在路径
	LogFolder() string
	// MiddlewareFolder 定义业务自己定义的中间件
	MiddlewareFolder() string
	// CommandFolder 定义业务定义的命令
	CommandFolder() string
	// RuntimeFolder 定义业务的运行中间态信息
	RuntimeFolder() string

	// LoadAppConfig 加载新的AppConfig，key为对应的函数转为小写下划线，比如ConfigFolder => config_folder
	LoadAppConfig(kv map[string]string)
}
```
