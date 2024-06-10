---
lang: zh-CN
title: 自动部署
description: 
---

# 自动部署

相关的命令详见：[deploy](../command/deploy)

部署自动化其实不是一个框架的刚需，有很多方式可以将一个服务进行自动化部署，比如现在比较流行的 Docker 化或者 CI/CD 流程。

但是一些比较个人比较小的项目，比如一个博客、一个官网网站，这些部署流程往往都太庞大了，更需要一个服务，能快速将在开发机器上写好、调试好的程序上传到目标服务器，并且更新应用程序。

这就是gob框架实现的发布自动化。

## SSH

所有的部署自动化工具，基本都依赖本地与远端服务器的连接，这个连接可以是 FTP，可以是 HTTP，但是更经常的连接是 SSH 连接。

基本上，SSH 账号是我们拿到 Web 服务器的首要凭证，所以要设计的自动化发布系统也是依赖 SSH 的。

对应的配置文件如下 `config/dev/ssh.yaml`，你可以看看每个配置的说明：

```yaml
timeout: 1s
network: tcp
host: 192.168.159.128
port: 22
username: demo
web-pwd:
  password: "123456"
web-key:
  rsa_key: "C:/Users/.ssh/id_rsa_manjarovm_demo_key"
  known_hosts: "C:/Users/.ssh/known_hosts"
web-ubuntu:
  host: 192.168.159.129
  username: cbh
  password: "5233"
```

SSH 的连接方式有两种，一种是直接使用用户名密码来连接远程服务器，还有一种是使用 rsa key 文件来连接远端服务器，所以这里的配置需要同时支持两种配置。

对于使用 rsa key 文件的方式，需要设置 rsk_key 的私钥地址和负责安全验证的 known_hosts。

## deploy

我们的 gob 框架是同时支持前后端的开发框架，所以自动化部署是需要同时支持前后端部署的，也就是说它的命令也需要支持前后端的部署，

```shell
./gob deploy frontend ，部署前端
./gob deploy backend ，部署后端
./gob deploy all ，同时部署前后端
./gob deploy rollback ，部署回滚
```

### 部署前端

你可以通过命令

```shell
./gob deploy frontend
```

或者 跳过编译环节

```shell
./gob deploy frontend -s=true
```

第一个方法会直接运行 npm run build，把前端代码生成在dist目录下，然后把dist目录下的文件上传到远端服务器，然后执行前置命令和后置命令。

而第二个方法会掉过编译，直接把dist目录下的文件上传到远端服务器，然后执行前置命令和后置命令。

### 部署后端

命令 `./gob deploy backend` 会自动编译gob二进制文件，然后上传到服务器上。

如果你的 post_action 设置的是重启远端服务器进程，那么实际上就是一个完整的cd行为了。

### 前后端一起部署

命令 `./gob deploy all`

### 部署回滚

每次部署执行，都会在本地的 deploy 目录下创建一个目录，目录名为当前时间戳，比如`20221214203745`。

如果你想回滚到上一次部署的版本，可以执行命令 `./gob deploy rollback 20221214203745 backend`。

实际上做的事情就是将 deploy 目录下的时间戳对应的文件再进行一次发布。