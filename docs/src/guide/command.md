---
lang: zh-CN
title: 命令
description: 
---

# 命令

## 指南

相关的命令详见：[command](../command/cmd.md)

gob 允许自定义命令，挂载到 gob 上。并且提供了`./gob command` 系列命令。

```shell
> ./gob command
控制台命令相关

Usage:
  gob command [flags]
  gob command [command]

Available Commands:
  list        列出所有控制台命令
  new         创建一个控制台命令

Flags:
  -h, --help   help for command

Use "gob command [command] --help" for more information about a command.

```

## 创建

创建一个新命令，可以使用 `./gob command new`

这是一个交互式的命令行工具。

创建完成之后，会在应用的 app/console/command/ 目录下创建一个新的文件。

## 自定义

gob 中的命令使用的是 cobra 库。 https://github.com/spf13/cobra

```
package command

import (
        "fmt"

        "github.com/chenbihao/gob/framework/cobra"
        "github.com/chenbihao/gob/framework/command/util"
)

var TestCommand = &cobra.Command{
        Use:   "test",
        Short: "test",
        RunE: func(c *cobra.Command, args []string) error {
                container := util.GetContainer(c.Root())
                fmt.Println(container)
                return nil
        },
}

```

基本上，我们要求实现

- Use // 命令行的关键字
- Short // 命令行的简短说明
- RunE // 命令行实际运行的程序

更多配置和参数可以参考 [cobra 的 github 页面](https://github.com/spf13/cobra)

## 挂载

编写完自定义命令后，请记得挂载到 `console/kernel.go` 中。

``` golang
func RunCommand(container framework.Container) error {
	var rootCmd = &cobra.Command{
		Use:   "gob",
		Short: "main",
		Long:  "gob commands",
	}

	ctx := commandUtil.RegiestContainer(container, rootCmd)
	gobCommand.AddKernelCommands(rootCmd)
	rootCmd.AddCommand(command.DemoCommand)
	return rootCmd.ExecuteContext(ctx)
}

```