---
lang: zh-CN
title: gob:cache
description: 
---
# gob:cache

## 说明

gob:cache 是直接微框架提供缓存服务，目前支持的缓存驱动有两种：

- redis
- memory

通过配置文件 `config/[env]/cache.yaml` 可以配置缓存服务的驱动和参数，如下是一个配置示例：

```yaml
#driver: redis # 连接驱动
#host: 127.0.0.1 # ip地址
#port: 6379 # 端口
#db: 0 #db
#timeout: 10s # 连接超时
#read_timeout: 2s # 读超时
#write_timeout: 2s # 写超时
#
driver: memory # 连接驱动
```

## 使用方法

cache 服务提供丰富的接口，可以通过接口来操作缓存，如下是接口定义：

```go


```

