---
lang: zh-CN
title: 版本
description: 
---
# 版本

## 命令

相关的命令详见：[version](../command/version)

gob 提供了查询当前版本和获取最新版本日志的命令

## 查询当前版本

使用命令 `gob version`

```
> ./gob version
gob version: 1.0.0
```

## 获取最新的版本

使用命令 `gob version list`

```
> ./gob version list       
===============前置条件检测===============
gob源码从github.com中下载，正在检测到github.com的连接
github.Rate{Limit:60, Remaining:59, Reset:github.Timestamp{2024-06-10 19:12:45 +0800 CST}}
gob源码从github.com中下载，github.com的连接正常
===============前置条件检测结束===============

最新的1个版本
-v0.1.11
  发布时间：2024-06-10 08:56:12
  修改说明：
    集成初始化脚手架，可通过以下命令在本地构建应用：

    使用 go install github.com/chenbihao/gob@latest 来安装 gob 命令。

    运行初始化脚手架 gob new 并根据命令行互动输入对应的应用名与模块名。

    进入对应的文件夹，使用 go mod tidy 安装相关依赖，
    随后可以通过引用 github.com/chenbihao/gob/framework 来引用框架相关模块

更多历史版本请参考 https://github.com/chenbihao/gob/releases

```