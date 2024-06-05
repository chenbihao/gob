---
lang: zh-CN
title: gob:distributed
description: 
---
# gob:distributed

## 说明

gob:distributed 是提供分布式选举的服务，可以用于分布式锁，分布式任务调度等场景。

当分布式集群中有需要选举出一个节点来执行任务时，可以使用此服务。

目前仅支持本地多进程的文件实现，后续会支持 redis 等分布式存储。

## 使用方法

```go

```