---
lang: zh-CN
title: gob:sls
description: 
---
# gob:sls

## 说明

sls 提供对接阿里云日志服务的服务，可以用于日志的收集。

## 配置

配置信息在 `config/[env]/app.yaml` 中配置，如下是一个配置示例：

```yaml
sls: # 阿里云SLS服务
    endpoint: cn-shanghai.log.aliyuncs.com
    access_key_id: your_access_key_id
    access_key_secret: your_access_key_secret
    project: gob
    logstore: gob_logstore
```

## 使用方法

```go

```