---
lang: zh-CN
title: 调试模式命令
description:
---
# 调试模式命令

## 命令介绍：
dev 调试工具，通过反向代理来管理前后端
## 前置需求：
前后端一体、cmd_build、config、app
## 支持命令：
```sh
./gob dev 			// 显示帮助信息
./gob dev frontend 	// 调试前端
./gob dev backend  	// 调试后端
./gob dev all  		// 显示所有
```
## 支持配置：
`app.yaml` 支持配置：
```yaml
dev: # 调试模式
port: 8070 # 调试模式最终监听的端口，默认为 8070
  frontend: # 前端调试模式配置
	port: 8071 # 前端监听端口, 默认 8071
  backend: # 后端调试模式配置
	refresh_time: 3  # 调试模式后端更新时间，如果文件变更，等待3s才进行一次更新，能让频繁保存变更更为顺畅, 默认1s
	port: 8072 # 后端监听端口，默认 8072
	monitor_folder: "" # 监听文件夹地址，为空或者不填默认为 AppFolder
```

## 使用方法：
稍后补全，可以使用`gob [command] help`命令获取相关帮助
