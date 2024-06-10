---
lang: zh-CN
title: 调试模式
description: 
---

# 调试模式

## 命令

相关的命令详见：[dev](../command/dev)

gob 框架自带调试模式，不管是前端还是后端，都可以启动调试模式，边修改代码，边编译运行服务。

对应的命令为 `./gob dev`

```shell
> ./gob dev
调试模式

Usage:
  gob dev [flags]
  gob dev [command]

Available Commands:
  all         同时启动前端和后端调试
  backend     启动后端调试模式
  frontend    启动前端调试模式

Flags:
  -h, --help   help for dev

Use "gob dev [command] --help" for more information about a command.

```

## 调试前端

使用命令 `./gob dev frontend`

要求当前编译机器安装 npm 软件，并且当前项目已经运行了 npm install，安装完成前端依赖。

实际上是调用 `npm run dev` 来调试前端

## 调试后端

使用命令 `./gob dev backend`

要求当前编译机器安装 go 软件，版本 > 1.3。

```shell
> ./gob dev backend
监控文件夹： path\gob\app
=============  后端编译开始 ============
build success please run ./gob direct
=============  后端编译成功 ============
启动后端服务:  http://127.0.0.1:8072
后端服务pid: 16092
代理服务启动: http://127.0.0.1:8070
[PID] 16092
app serve url :8072
```

> 后端调试默认是最后一次操作后3秒启动后端编译启动命令。
> gob 也允许通过配置修改这个等待时间。

# 同时调试

也可以选择同时调试，这个时候会同时运行调试前端和调试后端的程序

```shell
> ./gob dev all
```
