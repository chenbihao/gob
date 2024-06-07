package command

import (
	"fmt"
	"github.com/chenbihao/gob/framework/cobra"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/swaggo/swag/gen"
	"path/filepath"
)

/*
## 命令介绍：
swagger 生成
## 前置需求：
app
## 支持命令：
```sh
./gob swagger  			// 打印帮助信息
./gob swagger gen  		// 生成swagger文件
```
## 支持配置：无
*/

// 用于生成文档定位说明
const SwaggerCommandKey = "swagger命令"

func initSwaggerCommand() *cobra.Command {
	swaggerCommand.AddCommand(swaggerGenCommand)
	return swaggerCommand
}

var swaggerCommand = &cobra.Command{
	Use:   "swagger",
	Short: "swagger对应命令",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			c.Help()
		}
		return nil
	},
}

// swaggerGenCommand 生成具体的swagger文档
var swaggerGenCommand = &cobra.Command{
	Use:   "gen",
	Short: "生成对应的swagger文件, contain swagger.yaml, doc.go",
	Run: func(c *cobra.Command, args []string) {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		httpFolder := appService.HttpFolder()
		outputDir := filepath.Join(httpFolder, "swagger")

		conf := &gen.Config{
			// 遍历需要查询注释的目录
			SearchDir: httpFolder,
			// 不包含哪些文件
			Excludes: "",
			// 输出目录
			OutputDir: outputDir,
			// 输出类型 （这里应该有默认值的，但是运行时没找到？）
			OutputTypes: []string{"go", "json", "yaml"},
			// 整个swagger接口的说明文档注释
			MainAPIFile: "swagger.go",
			// 名字的显示策略，比如首字母大写等
			PropNamingStrategy: "",
			// 是否要解析vendor目录
			ParseVendor: false,
			// 是否要解析外部依赖库的包（ 0 none, 1 models, 2 operations, 3 all）
			ParseDependency: 0,
			// 是否要解析标准库的包
			ParseInternal: false,
			// 是否要查找markdown文件，这个markdown文件能用来为tag增加说明格式
			MarkdownFilesDir: httpFolder,
			// 是否应该在docs.go中生成时间戳
			GeneratedTime: false,
		}
		if err := gen.New().Build(conf); err != nil {
			fmt.Println(err)
		}
	},
}
