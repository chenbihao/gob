package command

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/chenbihao/gob/framework/cobra"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/util"
	"github.com/disiqueira/gotree"
	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"text/template"
)

/*
## 命令介绍：
command 命令
## 前置需求：
app
## 支持命令：
```sh
./gob command  		// 打印帮助信息
./gob command list  // 列出所有控制台命令
./gob command new  	// 创建一个控制台命令
```
## 支持配置：无
*/

// 用于生成文档定位说明
const CmdCommandKey = "命令"

// 初始化command相关命令
func initCmdCommand() *cobra.Command {
	cmdCommand.AddCommand(cmdListCommand)
	cmdCommand.AddCommand(cmdCreateCommand)
	return cmdCommand
}

// 二级命令
var cmdCommand = &cobra.Command{
	Use:   "command",
	Short: "控制台命令相关",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			c.Help()
		}
		return nil
	},
}

// cmdListCommand 列出所有的控制台命令
var cmdListCommand = &cobra.Command{
	Use:   "list",
	Short: "列出所有控制台命令",
	RunE: func(c *cobra.Command, args []string) error {
		rootTree := gotree.New("gob")
		addCommandToTree(rootTree, c.Root())
		fmt.Println(rootTree.Print())
		return nil
	},
}

// 添加命令到tree
func addCommandToTree(tree gotree.Tree, cmd *cobra.Command) {
	for _, c := range cmd.Commands() {
		node := tree.Add(c.Name() + "\t" + c.Short)
		addCommandToTree(node, c)
	}
}

// cmdCreateCommand 创建一个业务控制台命令
var cmdCreateCommand = &cobra.Command{
	Use:     "new",
	Aliases: []string{"create", "init"}, // 设置别名为 create init
	Short:   "创建一个控制台命令",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()

		fmt.Println("开始创建控制台命令...")
		var name string
		var folder string
		{
			prompt := &survey.Input{
				Message: "请输入控制台命令名称:",
			}
			err := survey.AskOne(prompt, &name)
			if err != nil {
				return err
			}
		}
		{
			prompt := &survey.Input{
				Message: "请输入文件夹名称(默认: 同控制台命令):",
			}
			err := survey.AskOne(prompt, &folder)
			if err != nil {
				return err
			}
		}

		if folder == "" {
			folder = name
		}

		// 判断文件不存在
		app := container.MustMake(contract.AppKey).(contract.App)

		pFolder := app.CommandFolder()
		subFolders, err := util.SubDir(pFolder)
		if err != nil {
			return err
		}
		subColl := util.NewCollection[string](subFolders)
		if subColl.Contains(folder) {
			fmt.Println("目录名称已经存在")
			return nil
		}

		// 开始创建文件
		if err := os.Mkdir(filepath.Join(pFolder, folder), 0700); err != nil {
			return err
		}

		// 创建title这个模版方法
		funcs := template.FuncMap{"title": cases.Title(language.Und, cases.NoLower).String}
		{
			//  创建name.go
			file := filepath.Join(pFolder, folder, name+".go")
			f, err := os.Create(file)
			if err != nil {
				return errors.Cause(err)
			}

			// 使用contractTmp模版来初始化template，并且让这个模版支持title方法，即支持{{.|title}}
			t := template.Must(template.New("cmd").Funcs(funcs).Parse(cmdTmpl))
			// 将name传递进入到template中渲染，并且输出到contract.go 中
			if err := t.Execute(f, name); err != nil {
				return errors.Cause(err)
			}
		}

		fmt.Println("创建新命令行工具成功，路径:", filepath.Join(pFolder, folder))
		fmt.Println("请记得开发完成后将命令行工具挂载到 console/kernel.go")
		return nil
	},
}

// 命令行工具模版
var cmdTmpl = `package {{.}}

import (
	"fmt"

	"github.com/chenbihao/gob/framework/cobra"
)

var {{.|title}}Command = &cobra.Command{
	Use:   "{{.}}",
	Short: "{{.}}",
	RunE: func(c *cobra.Command, args []string) error {
        container := c.GetContainer()
		fmt.Println(container)
		return nil
	},
}

`
