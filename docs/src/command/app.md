---
lang: zh-CN
title: 运行命令
description:
---
# 运行命令

## 命令介绍：
web app 业务应用控制命令
## 前置需求：无
## 支持命令：
```sh
./gob app start		// 启动一个 app 服务
	--address=:8080 // 指定端口
	--daemon=true 	// 守护模式（win下不支持）
./gob app state 	// 获取启动的 app 的信息
./gob app stop 		// 停止已经启动的 app 服务
./gob app restart 	// 重新启动一个 app 服务
```
## 支持配置：
```
app.address 		// 地址格式需符合 http.Server 的 Addr 格式
app.close_wait 		// 优雅关闭超时时间
```
## 支持环境变量：
```
ADDRESS				// 地址格式需符合 http.Server 的 Addr 格式
```

## 使用方法：
稍后补全，可以使用`gob [command] help`命令获取相关帮助
