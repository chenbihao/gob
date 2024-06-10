---
lang: zh-CN
title: 定时
description: 
---

# 定时任务

## 指南

相关的命令详见：[cron](../command/cron)

gob 中的定时任务是以命令的形式存在。gob 中也定义了一个命令 `./gob cron` 来对定时任务服务进行管理。

```shell
> ./gob cron
定时任务相关命令

Usage:
  gob cron [flags]
  gob cron [command]

Available Commands:
  list        列出所有的定时任务
  restart     重启cron常驻进程
  start       启动cron常驻进程
  state       cron常驻进程状态
  stop        停止cron常驻进程

Flags:
  -h, --help   help for cron

Use "gob cron [command] --help" for more information about a command.

```

# 创建

创建一个定时任务和创建命令（command）是一致的。具体参考[command](/command)

# 挂载

和挂载命令稍微有些不同，使用的方法是 `AddCronCommand`

```
rootCmd.AddCronCommand("* * * * *", command.DemoCommand)
```

# 查询

查询哪些定时任务挂载在服务上，使用命令 `./gob cron list`

# 启动

使用命令 `./gob cron start` 启动一个定时服务

也可以通过 `./gob cron start -d` 使用 daemon 模式启动一个定时服务

定时服务的输出记录在 `/storage/log/cron.log`

进程 id 记录在 `/storage/pid/app.pid`

# 状态

使用 daemon 模式启动定时服务的时候，可以使用命令 `./gob cron state` 查询定时任务状态

# 停止

使用 daemon 模式启动定时服务的时候，可以使用命令 `./gob cron stop` 停止定时任务

# 重启

使用 daemon 模式启动定时服务的时候，可以使用命令 `./gob cron restart` 重启定时任务


> 如果程序还未启动，调用 restart 命令，效果和 start 命令一样，daemon 模式启动定时服务
