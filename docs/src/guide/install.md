---
lang: zh-CN
title: 安装
description: 
---

# 安装

## 可执行文件

选择任一方式进行使用即可

### go install

已集成初始化脚手架，可通过以下命令在本地构建应用：

使用 `go install github.com/chenbihao/gob@latest` 来安装 gob 命令。

### 源码编译

clone git 地址：`git@github.com/chenbihao/gob.git`

在 gob 目录中运行命令 `go run main.go build self`

将生成的可执行文件 gob 放到 $PATH 目录中

### 下载源码

下载地址： [Releases](https://github.com/chenbihao/gob/releases)

## 初始化项目

使用命令 `gob new` 在当前目录创建子项目

这个创建新的 gob 项目的整个过程是交互式的。你可以根据命令行提示一步步进行:

```shell
> gob new
? 请输入目录名称： gobdemo
? 请输入模块名称(go.mod中的module, 默认为文件夹名称)：
? 请输入版本名称(参考 https://github.com/chenbihao/gob/releases，默认为最新版本)：
====================================================
开始进行创建应用操作
创建目录： D:\DevProjects\gobdemo
应用名称： gobdemo
gob框架版本： v0.1
创建临时目录 D:\DevProjects\template-gob-v0.1
...
```

创建项目成功后，会在执行目录下创建一个子目录，里面是 gob 的指定版本的代码。可以直接开始 gob 之旅。

使用 `go mod tidy` 下载依赖

接下来，可以通过命令 `go run main.go` 运行项目。

```shell
> go run main.go
gob 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令

Usage:
  gob [flags]
  gob [command]

Available Commands:
  app         业务应用控制命令
  build       编译相关命令
  command     控制台命令相关
  ...

Flags:
  -h, --help   help for gob

Use "gob [command] --help" for more information about a command.
```

至此，项目安装成功。
