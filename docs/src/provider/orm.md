---
lang: zh-CN
title: gob:orm
description:
---
# gob:orm

## 服务介绍：
提供ORM服务的服务，可以用于获取数据库连接，获取表结构等。
## 支持命令：无
## 支持配置：

使用之前需要确保已经正确配置了redis服务。

配置文件为 `config/[env]/database.yaml`，以下是一个配置的例子：

```yaml
##### mysql连接配置
#mysql:
#  hostname: 127.0.0.1
#  username: test
#  password: env(DB_PASSWORD)
#  timeout: 1

##### 分组下通用配置

conn_max_idle: 10 # 通用配置，连接池最大空闲连接数
conn_max_open: 100 # 通用配置，连接池最大连接数
conn_max_lifetime: 1h # 通用配置，连接数最大生命周期
protocol: tcp # 通用配置，传输协议
loc: Local # 通用配置，时区

##### 默认分组下的mysql连接配置
default:
  driver: mysql # 连接驱动
  dsn: "" # dsn，如果设置了dsn, 以下的所有设置都不生效
  host: localhost # ip地址
  port: 3306 # 端口
  database: demo # 数据库
  username: demo # 用户名
  password: "123456" # 密码
  allow_native_passwords: true
  charset: utf8mb4 # 字符集
  collation: utf8mb4_unicode_ci # 字符序
  timeout: 10s # 连接超时
  read_timeout: 2s # 读超时
  write_timeout: 2s # 写超时
  parse_time: true # 是否解析时间
  protocol: tcp # 传输协议
  loc: Local # 时区
  conn_max_idle: 10 # 连接池最大空闲连接数
  conn_max_open: 20 # 连接池最大连接数
  conn_max_lifetime: 1h # 连接数最大生命周期

##### 默认分组下的sqlite连接配置
#default:
#  driver: sqlite # 连接驱动
#  dsn: D:\dev-project\0.demo\out\box.db
```

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
