---
lang: zh-CN
title: swagger
description: 
---
# swagger

## 命令


相关的命令详见：[swagger](../command/swagger)

gob 使用 [swaggo](https://github.com/swaggo/swag) 集成了 swagger 生成和服务项目。

并且封装了 `./gob swagger` 命令。

```shell
> ./gob swagger
swagger对应命令

Usage:
  gob swagger [flags]
  gob swagger [command]

Available Commands:
  gen         生成对应的swagger文件, contain swagger.yaml, doc.go

Flags:
  -h, --help   help for swagger

Use "gob swagger [command] --help" for more information about a command.
```

## 注释

gob 使用 [swaggo](https://github.com/swaggo/swag) 来实现注释生成 swagger 功能。

全局注释在文件  `app/http/swagger.go` 中:

```go
// Package http API.
// @title gob
// @version 0.1.11
// @description gob框架
// @termsOfService https://github.com/swaggo/swag

// @contact.name chenbihao
// @contact.email chenbihao@foxmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}

package http

```

接口注释请写在各自模块的 `api.go` 中

```golang
// Demo godoc
// @Summary 获取所有用户
// @Description 获取所有用户
// @Produce  json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /demo/demo [get]
func (api *DemoApi) Demo(c *gin.Context) {
	users := api.service.GetUsers()
	usersDTO := UserModelsToUserDTOs(users)
	c.JSON(200, usersDTO)
}
```

swagger 注释的格式和关键词可以参考：[swaggo](https://github.com/swaggo/swag)

## 生成

使用命令 `./gob swagger gen`

```
> ./gob swagger gen
2024/06/10 18:35:22 Generate swagger docs....
2024/06/10 18:35:22 Generate general API Info, search dir:D:\DevProjects\自己库\gob\app\http
2024/06/10 18:35:23 Generating demo.UserDTO
2024/06/10 18:35:23 create docs.go at D:\DevProjects\自己库\gob\app\http\swagger/docs.go
2024/06/10 18:35:23 create swagger.json at D:\DevProjects\自己库\gob\app\http\swagger/swagger.json
2024/06/10 18:35:23 create swagger.yaml at D:\DevProjects\自己库\gob\app\http\swagger/swagger.yaml

```

在目录 `app/http/swagger/` 下自动生成swagger相关文件。

## 服务

可以使用命令 `./gob swagger serve` 启动当前应用的 swagger ui 服务。


> 如果你的 swagger 服务已经启动，更新 swagger 只需要重新运行 `./gob swagger gen` 就能更新。
> 
> 因为 swagger 服务读取的是生成的 `swagger.json` 这个文件。


服务端口，我们也可以通过配置文件 `config/[env]/swagger.yaml` 中的配置来配置swagger serve 启动的服务:

```
url: http://127.0.0.1:8069
```

