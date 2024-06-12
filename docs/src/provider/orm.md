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
	// 获取 DB
	GetDB(option ...DBOption) (*gorm.DB, error)

	// CanConnect 是否可以连接
	CanConnect(ctx context.Context, db *gorm.DB) (bool, error)

	// Table 相关
	GetTables(ctx context.Context, db *gorm.DB) ([]string, error)
	HasTable(ctx context.Context, db *gorm.DB, table string) (bool, error)
	GetTableColumns(ctx context.Context, db *gorm.DB, table string) ([]TableColumn, error)
}
```
