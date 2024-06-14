package model

import (
	"github.com/chenbihao/gob/framework/cobra"
)

/*
## 命令介绍：
数据库模型相关的命令
## 前置需求：
## 支持命令：
```sh
./gob model test 	// 测试数据库连接，并展示数据表列表

./gob model gen  	// 通过数据库生成模型
	-d, --database string   // 可选，模型连接的数据库 (default "database.default")
	-t, --table string      // 可选，模型连接的数据表
	-o, --output string     // 必填，模型输出地址

./gob model api  	// 通过数据库生成api
	-d, --database string   // 可选，模型连接的数据库 (default "database.default")
	-t, --table string      // 可选，模型连接的数据表
	-o, --output string     // 必填，模型输出地址

```
## 支持配置：
无
## 支持环境变量：
无
*/

// 用于生成文档定位说明
const ModelCommandKey = "生成命令"

// 代表输出路径
var output string

// 代表数据库连接
var database string

// 代表表格
var table string

// InitModelCommand 获取model相关的命令
func InitModelCommand() *cobra.Command {

	// model test
	modelTestCommand.Flags().StringVarP(&database, "database", "d", "database.default", "模型连接的数据库")
	modelCommand.AddCommand(modelTestCommand)

	// model gen
	modelGenCommand.Flags().StringVarP(&database, "database", "d", "database.default", "模型连接的数据库")
	modelGenCommand.Flags().StringVarP(&table, "table", "t", "", "模型连接的数据表")
	modelGenCommand.Flags().StringVarP(&output, "output", "o", "", "模型输出地址")
	_ = modelGenCommand.MarkFlagRequired("output")
	modelCommand.AddCommand(modelGenCommand)

	// model api
	modelApiCommand.Flags().StringVarP(&database, "database", "d", "database.default", "模型连接的数据库")
	modelApiCommand.Flags().StringVarP(&table, "table", "t", "", "模型连接的数据表")
	modelApiCommand.Flags().StringVarP(&output, "output", "o", "", "模型输出地址, 文件夹地址")
	_ = modelApiCommand.MarkFlagRequired("output")
	modelCommand.AddCommand(modelApiCommand)
	return modelCommand
}

// modelCommand 模型相关的命令
var modelCommand = &cobra.Command{
	Use:   "model",
	Short: "数据库模型相关的命令",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			c.Help()
		}
		return nil
	},
}
