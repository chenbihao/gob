---
lang: zh-CN
title: 待办
description:
---

# 待办

## 新版设计

移除默认web框架，默认是启动命令框架，自主选择接入 web/wailsApp 等

框架本体：快速开发框架，通过运行gob命令，快速生成项目骨架，快速开发功能

需要集成：

- 默认绑定gob自带的功能 gob\app\config\log\version\new\build？
- command\provider\model\middleware 其他按需挂载
- 默认绑定gobRootCommand，可完全替换也可继续使用，允许传入新的rootCommand名称

- 契约改造联动？ 查询依赖是否已经全bind、挂载命令、挂载目录、挂载配置？
- 文档自动化：依赖、支持命令、支持配置

### 框架自带功能

- 框架定义 App
- 命令行框架 Cobra v1.9.1

## Todo:

### 改变清单：

- [ ] 配置：  [koanf](https://github.com/knadh/koanf) 模块化且支持pflag
- [ ] 日志：      `log/slog`
- [ ] 日志切割：  `lumberjack.v3`? / `go-rotate-logs` 替换掉 `file-rotatelogs`
- [ ] 命令行交互：`bubbletea`   替换掉 `survey`
- [ ] 新增：
  - `samber/lo` Go 1.18+ 泛型的实用程序库

### 配置重构

- [x] 移除原config逻辑：
  1. app - 先读取 sysEnv 与 args 和空 configMap （先设置一遍Folder相关
  2. config - 配置初传递 appService envService
  3. config - 读取每个文件配置 gobConf.loadConfigFile
      1. 替换源文本 env(key)
      2. 单独更新 app.path 里的配置 （回调app.LoadAppConfig）
      3. 监听刷新配置 写入or删除
  4. find -> searchMap -> 迭代遍历key

新config逻辑：初始化 koanf，其他模块需要则填充进去并刷新获取

- [x] 读取环境变量（先读取环境变量，后读取.env文件，可以替换值）
- [x] 初始化 
  - 默认是极简模式（根目录直接config.yaml）
  - 可选开启配置文件夹
  - 可选开启ENV（deploy_env：env/test/prod）

- [ ] 配置结构化
- [ ] 其他子配置单独管理，开放挂载配置入参，挂载则刷新配置
- [ ] 读写锁

### 挂载与蓝图

- [ ] 命令行入参？
- [ ] 挂载命令

- [ ] app : Folder 结构化？ 挂载目录？
- [ ] 配置分离 加个简单加密存储的配置文件 SQLite ？

- [ ] 初始化蓝图流程
  - [ ] 蓝图定义，包括依赖关系、版本等
  - [ ] 定义拉取蓝图模块流程
  - [ ] 拉取后执行表迁移工作


### 框架模块优化

- [ ] 发包时不包含docs等其他内容，这样拉取的时候不用清理

- [ ] 日志
    - [ ] 同时支持多个日志输出
    - [ ] 接管gin的日志输出
    - [ ] 统一优化日志打印格式
    - [ ] 可选固定 json 字段顺序打印配置（强迫症可选）
    - [ ] 日志库切换？
        - [ ] file-rotatelogs 废弃

- [ ] 数据库重连重试机制
- [ ] model 代码生成功能优化

    - [ ] 生成更符合业务调用场景
    - [ ] 其他数据源如sqlite的字段优化

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
    - [ ] gspt 构建出错，需要交叉编译

- [ ] 前端文件夹配置可选？
    - [ ] frontendFolder 配置可选

- [ ] 其他优化
    - [ ] 把其他命令适配到纯工具模式（`go install`）
    - [ ] win 下不支持 Daemon，兼容成后台运行（appDaemon）
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
- [ ] 目前考虑 dbmate
    - 需要重新思考model生成相关逻辑
    - 先手写model后表？
    - 先手写sql后model和api？
    - 都兼容？

- [ ] AI能力接入？
  - MCP服务获取框架文档和代码规范，方便AI实现代码

### 蓝图模块功能

- [ ] 后台管理基础
    - [ ] 低代码快速搭建？

- [ ] 用户注册登录
    - [ ] RBAC 权限
    - [ ] 多租户模块

- [ ] 博客
- [ ] 加密
- [ ] ...

## 其他功能优化

- [ ] 部分 linux 内容未测试
    - [ ] 条件编译
    - [ ] 守护进程模式 `app start --daemon=true`
    - [ ] gspt 库（`CGO_ENABLED=1`）

- [ ] 业务单测的构建

## 归档

### 梳理相关

- [X] 梳理三方库引入
    - fsnotify、go-daemon、goconvey、swaggo、cast
    - survey/v2、go-git/v5、go-github/v62、go-redis/v9、cron/v3、gorm + gen
    - gotree、uuid、xid、ratelimit、mapstructure
    - kr/pretty、jennifer/jen
    - 预计移除：jianfengye/collection
    - 预计移除：file-rotatelogs、natefinch/lumberjack
    - 预计新增：samber/lo
