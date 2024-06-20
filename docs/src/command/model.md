---
lang: zh-CN
title: 生成命令
description:
---
# 生成命令

## 命令介绍：
数据库模型相关的命令
## 前置需求：
## 支持命令：
```sh
./gob model test 	// 测试数据库连接，并展示数据表列表

./gob model gen  	// 通过数据库生成模型
	-d, --database string   // 可选，模型连接的数据库 (default "database.default")
	-t, --table string      // 可选，模型连接的数据表
	-o, --output string     // 必填，模型输出地址

./gob model api  	// 通过数据库生成api
	-d, --database string   // 可选，模型连接的数据库 (default "database.default")
	-t, --table string      // 可选，模型连接的数据表
	-o, --output string     // 必填，模型输出地址

```
## 支持配置：
需要配置数据库连接，具体查看：[orm](../provider/orm)
## 支持环境变量：无

## 使用方法：
稍后补全，可以使用`gob [command] help`命令获取相关帮助
