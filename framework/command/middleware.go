package command

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/chenbihao/gob/framework/cobra"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/util"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-git/go-git/v5"
)

/*
## 命令介绍：
middleware 生成
## 前置需求：
app
## 支持命令：
```sh
./gob middleware  			// 打印帮助信息
./gob middleware list  		// 列出所有的业务中间件
./gob middleware new  		// 创建一个新的业务中间件
./gob middleware migrate 	// 迁移 Gin 已有的中间件
```
## 支持配置：无
*/

// 用于生成文档定位说明
const MiddlewareCommandKey = "中间件命令"

// 初始化中间件相关命令
func initMiddlewareCommand() *cobra.Command {
	middlewareCommand.AddCommand(middlewareAllCommand)
	middlewareCommand.AddCommand(middlewareMigrateCommand)
	middlewareCommand.AddCommand(middlewareCreateCommand)
	return middlewareCommand
}

// middlewareCommand 中间件二级命令
var middlewareCommand = &cobra.Command{
	Use:   "middleware",
	Short: "中间件相关命令",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			c.Help()
		}
		return nil
	},
}

// middlewareAllCommand 显示所有安装的中间件
var middlewareAllCommand = &cobra.Command{
	Use:   "list",
	Short: "显示所有中间件",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		middlewarePath := path.Join(appService.BaseFolder(), "app", "http", "middleware")

		// 读取文件夹
		files, err := os.ReadDir(middlewarePath)
		if err != nil {
			return err
		}
		for _, f := range files {
			if f.IsDir() {
				fmt.Println(f.Name())
			}
		}
		return nil
	},
}

// 从gin-contrib中迁移中间件
var middlewareMigrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "迁移gin-contrib中间件, 迁移地址：https://github.com/gin-contrib/[middleware].git",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		fmt.Println("迁移一个Gin中间件")
		var repo string
		{
			prompt := &survey.Input{
				Message: "请输入中间件名称：",
			}
			err := survey.AskOne(prompt, &repo)
			if err != nil {
				return err
			}
		}
		// step2 : 下载git到一个目录中
		appService := container.MustMake(contract.AppKey).(contract.App)

		middlewarePath := appService.MiddlewareFolder()
		url := "https://github.com/gin-contrib/" + repo + ".git"
		fmt.Println("下载中间件 gin-contrib:")
		fmt.Println(url)
		_, err := git.PlainClone(path.Join(middlewarePath, repo), false, &git.CloneOptions{
			URL:      url,
			Progress: os.Stdout,
		})
		if err != nil {
			return err
		}

		// step3 : 删除不必要的文件如 go.mod, go.sum, .git 等
		repoFolder := path.Join(middlewarePath, repo)
		removeUnnecessaryFiles(repoFolder)

		// step4 : 替换关键词
		filepath.Walk(repoFolder, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if filepath.Ext(path) != ".go" {
				return nil
			}
			c, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			isContain := bytes.Contains(c, []byte("github.com/gin-gonic/gin"))
			if isContain {
				fmt.Println("更新文件:" + path)
				c = bytes.ReplaceAll(c, []byte("github.com/gin-gonic/gin"), []byte("github.com/chenbihao/gob/framework/gin"))
				err = os.WriteFile(path, c, 0644)
				if err != nil {
					return err
				}
			}

			return nil
		})
		return nil
	},
}

func removeUnnecessaryFiles(repoFolder string) {
	files := []string{"go.mod", "go.sum", ".golangci.yml", ".goreleaser.yaml", ".gitignore"}
	removeFiles(repoFolder, files)
	dirNames := []string{".git", ".github"}
	removeDir(repoFolder, dirNames)
}

func removeFiles(folder string, files []string) {
	for _, file := range files {
		filePath := path.Join(folder, file)
		fmt.Println("remove " + filePath)
		err := os.Remove(filePath)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func removeDir(repoFolder string, dirNames []string) {
	for _, dirName := range dirNames {
		dirPath := path.Join(repoFolder, dirName)
		fmt.Println("remove " + dirPath)
		err := os.RemoveAll(dirPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// providerCreateCommand 创建一个新的服务，包括服务提供者，服务接口协议，服务实例
var middlewareCreateCommand = &cobra.Command{
	Use:     "new",
	Aliases: []string{"create", "init"},
	Short:   "创建一个中间件",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		fmt.Println("创建一个中间件")
		var name string
		var folder string
		{
			prompt := &survey.Input{
				Message: "请输入中间件名称：",
			}
			err := survey.AskOne(prompt, &name)
			if err != nil {
				return err
			}
		}
		{
			prompt := &survey.Input{
				Message: "请输入中间件所在目录名称(默认: 同中间件名称):",
			}
			err := survey.AskOne(prompt, &folder)
			if err != nil {
				return err
			}
		}

		if folder == "" {
			folder = name
		}

		app := container.MustMake(contract.AppKey).(contract.App)

		pFolder := app.MiddlewareFolder()
		subFolders, err := util.SubDir(pFolder)
		if err != nil {
			return err
		}
		subColl := util.NewCollection[string](subFolders)
		if subColl.Contains(folder) {
			fmt.Println("目录已经存在")
			return nil
		}

		// 开始创建文件
		if err := os.Mkdir(filepath.Join(pFolder, folder), 0700); err != nil {
			return err
		}
		funcs := template.FuncMap{"title": util.ToTitle}
		{
			//  创建
			file := filepath.Join(pFolder, folder, "middleware.go")
			f, err := os.Create(file)
			if err != nil {
				return errors.Cause(err)
			}

			t := template.Must(template.New("middleware").Funcs(funcs).Parse(middlewareTmp))
			if err := t.Execute(f, name); err != nil {
				return errors.Cause(err)
			}
		}
		fmt.Println("创建中间件成功, 文件夹地址:", filepath.Join(pFolder, folder))
		return nil
	},
}

var middlewareTmp = `package {{.}}

import "github.com/chenbihao/gob/framework/gin"

// {{.|title}}Middleware 代表中间件函数
func {{.|title}}Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}

`
