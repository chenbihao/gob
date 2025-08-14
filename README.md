# Gob

基于 go 语言编写的快速开发框架，

## 官方网站

[详细文档](https://chenbihao.github.io/gob/)

## 框架特色



## 使用指南

已集成初始化脚手架，可通过以下命令在本地构建应用：

使用 `go install github.com/chenbihao/gob@latest` 来安装 gob 命令。

在目标文件夹，运行初始化脚手架 `gob new` 并根据命令行互动输入对应的应用名与模块名。

进入对应的文件夹，使用 `go mod tidy` 安装相关依赖，
随后可以通过引用 `github.com/chenbihao/gob/framework` 来引用框架相关模块

## 技术栈

具备快速开发框架的基础能力

 - go 1.24+
<!--
- gin v1.9.1
- gorm v1.25.9
- swagger
- vue3
- ,,, -->

## 服务提供者

提供了场景的功能封装提供，例如：

- 

详见 [文档-服务提供者](https://chenbihao.github.io/gob/provider/)

## 命令行工具

提供了提效命令工具，例如：

- 

详见 [文档-提供命令](https://chenbihao.github.io/gob/command/)

## 蓝图

Todo...

计划实现“蓝图”功能。

一个快速开发应用模板拉取的功能，框架将提供一些方便的“蓝图”，例如后台管理蓝图、权限蓝图、博客蓝图等。

用户拉取“蓝图”后，框架通过数据库版本管理能力，快速搭建起一个具备基础能力的服务。

## 计划

[Todo 列表](docs/src/guide/TODO.md)

## 更多

有任何问题可直接 github 留言，或者联系作者。

本框架是作者在学习手写开发框架后的产物，计划持续开发并作为开发项目用的脚手架。
