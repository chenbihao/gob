---
lang: zh-CN
title: 环境变量
description: 
---

# 环境变量

## 设置

gob 支持使用应用默认下的隐藏文件 `.env` 来配置各个机器不同的环境变量。

```
APP_ENV=dev

DB_PASSWORD=mypassword
```

环境变量的设置可以在配置文件中通过 `env([环境变量])` 来获取到。

比如：

```
mysql:
    hostname: 127.0.0.1
    username: root
    password: env(DB_PASSWORD)
    timeout: 1
    readtime: 1
```

## 应用环境

gob 启动应用的默认应用环境为 dev。

你可以通过设置 `.env` 文件中的 `APP_ENV` 设置应用环境。

应用环境选择：

- `dev`  // 开发使用
- `test` // 测试环境
- `prod` // 线上使用

应用环境对应配置的文件夹，配置服务会去对应应用环境的文件夹中寻找配置。

比如应用环境为 dev，在代码中使用

```go
configService := container.MustMake(contract.ConfigKey).(contract.Config)
url := configService.GetString("app.url")
```

查找文件为：`config/dev/app.yaml`

通过命令`./gob env`也可以获取当前应用环境：

```shell
> ./gob env
environment: dev
```

## 命令

相关的命令详见：[env](../command/env)