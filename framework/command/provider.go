package command

import (
	"fmt"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/cobra"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/util"
	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"text/template"

	// survey 于2023年9月24日 停止维护，新推荐 github.com/charmbracelet/bubbletea
	"github.com/AlecAivazis/survey/v2"
)

/**
	命令介绍：
		provider 生成
	前置需求：
		app
	支持命令：
		./gob provider  		// 打印帮助信息
		./gob provider list  	// 列出容器内的所有服务的字符串凭证
		./gob provider new  	// 创建一个服务
	支持配置：
		无
**/

// 初始化provider相关服务
func initProviderCommand() *cobra.Command {
	providerCommand.AddCommand(providerListCommand)
	providerCommand.AddCommand(providerCreateCommand)
	return providerCommand
}

// 二级命令
var providerCommand = &cobra.Command{
	Use:   "provider",
	Short: "服务提供相关命令",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			c.Help()
		}
		return nil
	},
}

// providerListCommand 列出容器内的所有服务
var providerListCommand = &cobra.Command{
	Use:   "list",
	Short: "列出容器内的所有服务",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		gobContainer := container.(*framework.GobContainer)
		// 获取字符串凭证
		list := gobContainer.NameList()
		// 打印
		for _, line := range list {
			println(line)
		}
		return nil
	},
}

// providerCreateCommand 创建一个新的服务，包括服务提供者，服务接口协议，服务实例
var providerCreateCommand = &cobra.Command{
	Use:     "new",
	Aliases: []string{"create", "init"},
	Short:   "创建一个服务",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		fmt.Println("创建一个服务")
		var name string
		var folder string
		{
			prompt := &survey.Input{
				Message: "请输入服务名称(服务凭证)：",
			}
			if err := survey.AskOne(prompt, &name); err != nil {
				return err
			}
		}
		{
			prompt := &survey.Input{
				Message: "请输入服务所在目录名称(默认: 同服务名称):",
			}
			if err := survey.AskOne(prompt, &folder); err != nil {
				return err
			}
		}

		// 检查服务是否存在
		providers := container.(*framework.GobContainer).NameList()
		providerColl := util.NewCollection[string](providers)
		if providerColl.Contains(name) {
			fmt.Println("服务名称已经存在")
			return nil
		}
		if folder == "" {
			folder = name
		}

		app := container.MustMake(contract.AppKey).(contract.App)

		pFolder := app.ProviderFolder()
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
			//  创建contract.go
			file := filepath.Join(pFolder, folder, "contract.go")
			f, err := os.Create(file)
			if err != nil {
				return errors.Cause(err)
			}

			// 使用contractTmp模版来初始化template，并且让这个模版支持title方法，即支持{{.|title}}
			t := template.Must(template.New("contract").Funcs(funcs).Parse(contractTmp))
			// 将name传递进入到template中渲染，并且输出到contract.go 中
			if err := t.Execute(f, name); err != nil {
				return errors.Cause(err)
			}
		}
		{
			// 创建provider.go
			file := filepath.Join(pFolder, folder, "provider.go")
			f, err := os.Create(file)
			if err != nil {
				return err
			}
			t := template.Must(template.New("provider").Funcs(funcs).Parse(providerTmp))
			if err := t.Execute(f, name); err != nil {
				return err
			}
		}
		{
			//  创建service.go
			file := filepath.Join(pFolder, folder, "service.go")
			f, err := os.Create(file)
			if err != nil {
				return err
			}
			t := template.Must(template.New("service").Funcs(funcs).Parse(serviceTmp))
			if err := t.Execute(f, name); err != nil {
				return err
			}
		}
		fmt.Println("创建服务成功, 文件夹地址:", filepath.Join(pFolder, folder))
		fmt.Println("请不要忘记挂载新创建的服务")
		return nil
	},
}

var contractTmp = `package {{.}}

const {{.|title}}Key = "{{.}}"

type Service interface {
	// 请在这里定义你的方法
	Foo() string
}
`

var providerTmp = `package {{.}}

// ServiceProvider 实现文件 provider.go

import (
	"github.com/chenbihao/gob/framework"
)

type {{.|title}}Provider struct {
	c framework.Container
}
var _ framework.ServiceProvider = (*{{.|title}}Provider)(nil)

func (sp *{{.|title}}Provider) Name() string {
	return contract.{{.|title}}Key
}

func (sp *{{.|title}}Provider) Register(c framework.Container) framework.NewInstance {
	return New{{.|title}}Service
}

func (sp *{{.|title}}Provider) IsDefer() bool {
	return false
}

func (sp *{{.|title}}Provider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *{{.|title}}Provider) Boot(c framework.Container) error {
	return nil
}

`

var serviceTmp = `package {{.}}

// 实现具体的服务实例 service.go

import "github.com/chenbihao/gob/framework"

type {{.|title}}Service struct {
	container framework.Container
}

func New{{.|title}}Service(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	return &{{.|title}}Service{container: container}, nil
}

`
