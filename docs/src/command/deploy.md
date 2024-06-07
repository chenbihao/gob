---
lang: zh-CN
title: 部署命令
description:
---
# 部署命令

## 命令介绍：
部署命令
## 前置需求：
app
## 支持命令：
```sh
./gob deploy frontend`	// 部署前端
	-s --skip-build     	// 跳过前端构建
./gob deploy backend`	// 部署后端
./gob deploy all`		// 同时部署前后端
	-s --skip-build     	// 跳过前端构建
./gob deploy rollback`	// 部署回滚
```
## 支持配置：
`deploy.yaml` 支持配置：
```yaml
connections: # 要自动化部署的连接
  - ssh.web-key

remote_folder: "/home/demo/deploy/"  # 远端的部署文件夹

frontend: # 前端部署配置
  pre_action: # 部署前置命令
	- "pwd"
  post_action: # 部署后置命令
	- "pwd"

backend: # 后端部署配置
  goos: linux # 部署目标操作系统
  goarch: amd64 # 部署目标cpu架构
  pre_action: # 部署前置命令
	- "rm /home/demo/deploy/gob"
  post_action: # 部署后置命令
	- "chmod 777 /home/demo/deploy/gob"
	- "/home/demo/deploy/gob app restart"
```
ssh 支持配置：详见 contract/redis.go

## 使用方法：

