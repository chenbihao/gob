---
lang: zh-CN
title: 待办
description: 
---
# 待办

## 功能优化

- [ ] 部分 linux 内容未测试
    - [ ] 条件编译
    - [ ] 守护进程模式 `app start --daemon=true`
    - [ ] gspt 库（CGO_ENABLED=1）


- [ ] 业务单测的构建

## 框架支持功能

### 框架模块优化

- [ ] 脚手架优化
    - [ ] 构建后读取配置问题（生产部署时如何读取配置问题）
    - [ ] `go install` 使用的优化（编译后的命令行执行问题）
    - [ ] 前端文件夹配置可选？
        - [ ] frontendFolder 配置 可选而不是写死
        - [ ] frontendFolder/dist 可选？

- [ ] 框架
    - [ ] 抽取配置等魔术编码，例如 `"app.pid"` 等
    - [ ] 文件创建相关（app 获取路径时，有些可以判断并创建）
    - [ ] 获取各类 Folder 时就校验并判断，而不是分散在 app 或者 cron 服务里判断
    - [ ] win 下不支持 Daemon，兼容成后台运行（appDaemon）

- [ ] 日志
    - [ ] 同时支持多个日志输出
    - [ ] 统一优化日志打印格式
    - [ ] 可选固定 json 字段顺序打印配置（强迫症可选）

- [ ] model 代码生成功能优化
    - [ ] 生成更符合业务调用场景
    - [ ] 其他数据源如sqlite的字段优化
    - [ ] 联动

- [ ] 调试模式
    - [ ] 文件监控比服务启动还提前，编译完成前修改文件可能会导致空指针
    - [ ] swagger 热更新？热读取版本配置？

- [ ] 部署模式
    - [ ] 当断开 ssh 时服务会停止（特定问题，Manjaro 系统特有）
    - [ ] 发布自动化完善（部署备份、部署回滚）

- [ ] 将过时或停止维护的三方库换掉
    - [ ] survey 换成 bubbletea
        - [ ] 命令行支持静默运行参数
    - [ ] github 的调用限制重构复用
    - [ ] gspt 构建出错 ， 需要交叉编译


- [ ] 数据库重连重试机制
- [ ] cache 服务当配置了 redis，并且有 redis 相关配置时，优先读取 redis 配置

### 框架模块新增

- [ ] 远程配置中心

- [ ] 引入数据库迁移，方便后续蓝图，选型：
  - 简单 gorm 迁移增强
    - [gormigrate](https://github.com/go-gormigrate/gormigrate)
  - 驱动多功能强大，无建表
    - [migrate](https://github.com/golang-migrate/migrate)
  - 驱动少功能强大，有建表，有导出架构
    - [dbmate](https://github.com/amacneil/dbmate)
  - 目前考虑 dbmate

### 蓝图模块功能

- [ ] 初始化蓝图流程
    - [ ] 定义拉取蓝图模块流程
    - [ ] 拉取后执行表迁移工作


- [ ] 后台管理基础
    - [ ] 低代码快速搭建？
- [ ] 用户注册登录
    - [ ] RBAC 权限
    - [ ] 多租户模块
- [ ] 博客
- [ ] ...
- [ ] ...



## 已完成归档

### 梳理相关

- [x] 梳理使用框架
    - [x] 梳理源码引入
        - cobra
        - gin v1.9.1 + middleware
    - [x] 梳理三方库引入
        - fsnotify、go-daemon、goconvey、swaggo、cast
        - survey/v2、go-git/v5、go-github/v62、go-redis/v9、cron/v3、gorm + gen
        - gotree、uuid、xid、ratelimit、file-rotatelogs、mapstructure
        - jennifer/jen、jianfengye/collection、kr/pretty
    - [x] 梳理三方框架使用
        - vue、vuepress

- [x] 梳理新版本 go 废弃 API，换成新的
    - `io/ioutil` -> `os`、`io`
    - `strings.Title` -> `cases.Title`
    - `math/rand` -> `rand.Rand`

### 统一代码

- [x] 统一 provider 注册方法 （`func (provider *GormProvider) Register` 里的调用 new 命名）
- [x] 补充 command 、contract 文件开头说明文档，方便查看（甚至改成支持 doc）
    - [x] command：包括命令说明、可选配置项
    - [x] contract：包括对应命令、配置项说明





