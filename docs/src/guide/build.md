---
lang: zh-CN
title: 编译
description: 
---
# 编译

## 命令

应用分为前端（frontend）和后端（backend），所以编译也分为三类
- 编译前端
- 编译后端
- 自编译
- 同时编译

相关的命令详见：[build](../command/build)

```
> ./gob build
编译相关命令

Usage:
  gob build [flags]
  gob build [command]

Available Commands:
  all         同时编译前端和后端
  backend     使用 go 编译后端
  frontend    使用 npm 编译前端
  self        编译 gob 命令

Flags:
  -h, --help   help for build

Use "gob build [command] --help" for more information about a command.

```

## 编译前端

要求当前编译机器安装 npm 软件，并且当前项目已经运行了 npm install，安装完成前端依赖。

运行命令 `./gob build frontend`

编译后的前端文件在 dist 目录中

实际上 build 就是调用 `npm build` 来编译前端项目。


## 编译后端

要求当前编译机器安装 go 软件，版本 > 1.3。

运行命令： `./gob build backend`

在项目根目录下就看到生成的可执行文件 gob。 后续可以通过 ./gob 直接运行。

## 自编译

在项目根目录下，gob 可以通过 gob 命令编译出 gob 命令自己。

运行命令 `gob build self`

在项目根目录下就看到生成的可执行文件 gob。 后续可以通过 ./gob 直接运行。

> 其实自编译和后端编译是同样效果，但是为了命令语义化，增加了自编译的命令。

## 同时编译

顾名思义，同时编译前端和后端，命令为 `./gob build all`