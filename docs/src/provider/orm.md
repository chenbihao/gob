---
lang: zh-CN
title: gob:orm
description:
---
# gob:orm

## 服务介绍：
提供ORM服务的服务，可以用于获取数据库连接，获取表结构等。
## 支持命令：无
## 支持配置：无

## 提供方法：
```go 
type ORM interface {
	GetDB(option ...DBOption) (*gorm.DB, error)
}
```
