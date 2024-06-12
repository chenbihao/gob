package command

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/chenbihao/gob/framework/cobra"
	"github.com/chenbihao/gob/framework/util"

	"github.com/AlecAivazis/survey/v2"
	"github.com/google/go-github/v58/github"
	"github.com/spf13/cast"
)

/*
## 命令介绍：
拉取最新的 gob 框架内容
## 前置需求：无
## 支持命令：
```sh
./gob new		// 拉取框架并创建一个新的应用
```
## 支持配置：无
*/

// 用于生成文档定位说明
const NewCommandKey = "创建命令"

// new相关的名称
func initNewCommand() *cobra.Command {
	return newCommand
}

// 创建一个新应用
var newCommand = &cobra.Command{
	Use:     "new",
	Aliases: []string{"create", "init"},
	Short:   "创建一个新的应用",
	RunE: func(c *cobra.Command, args []string) error {
		currentPath := util.GetExecDirectory()

		var name string
		var folder string
		var mod string
		var version string
		var release *github.RepositoryRelease
		{
			prompt := &survey.Input{
				Message: "请输入目录名称：",
			}
			err := survey.AskOne(prompt, &name)
			if err != nil {
				fmt.Println("任务终止：" + err.Error())
				return err
			}

			folder = filepath.Join(currentPath, name)
			if util.Exists(folder) {
				isForce := false
				prompt2 := &survey.Confirm{
					Message: "目录" + folder + "已经存在,是否删除重新创建？(确认后立刻执行删除操作！)",
					Default: false,
				}
				err := survey.AskOne(prompt2, &isForce)
				if err != nil {
					fmt.Println("任务终止：" + err.Error())
					return err
				}

				if isForce {
					if err := os.RemoveAll(folder); err != nil {
						fmt.Println("任务终止：" + err.Error())
						return err
					}
				} else {
					fmt.Println("目录已存在，创建应用失败")
					return nil
				}
			}
		}
		{
			prompt := &survey.Input{
				Message: "请输入模块名称(go.mod中的module, 默认为文件夹名称)：",
			}
			err := survey.AskOne(prompt, &mod)
			if err != nil {
				fmt.Println("任务终止：" + err.Error())
				return err
			}
			if mod == "" {
				mod = name
			}
		}
		{
			// 获取gob的版本
			// 检测到github的连接
			fmt.Println("gob源码从github.com中下载，正在检测到github.com的连接")
			var client *github.Client
			client = github.NewClient(nil)
			perPage := 10
			opts := &github.ListOptions{Page: 1, PerPage: perPage}
			releases, rsp, err := client.Repositories.ListReleases(context.Background(), Owner, Repo, opts)
			fmt.Println(rsp.Rate.String())
			if err != nil {
				if _, ok := err.(*github.RateLimitError); ok {
					fmt.Println("错误提示：" + err.Error())
					fmt.Println("说明你的出口ip遇到github的调用限制，可以使用github.com帐号登录方式来增加调用次数")
					githubUserName := ""
					prompt := &survey.Input{
						Message: "请输入github帐号用户名：",
					}
					if err := survey.AskOne(prompt, &githubUserName); err != nil {
						fmt.Println("任务终止：" + err.Error())
						return nil
					}
					githubPassword := ""
					promptPwd := &survey.Password{
						Message: "请输入github帐号密码：",
					}
					if err := survey.AskOne(promptPwd, &githubPassword); err != nil {
						fmt.Println("任务终止：" + err.Error())
						return nil
					}

					httpClient := &http.Client{
						Transport: &http.Transport{
							Proxy: func(req *http.Request) (*url.URL, error) {
								req.SetBasicAuth(githubUserName, githubPassword)
								return nil, nil
							},
						},
					}
					client = github.NewClient(httpClient)
					releases, rsp, err = client.Repositories.ListReleases(context.Background(), Owner, Repo, opts)
					if err != nil {
						fmt.Println("错误提示：" + err.Error())
						fmt.Println("用户名密码错误，请重新开始")
						return nil
					}
					if len(releases) == 0 {
						fmt.Println("用户名密码错误，请重新开始")
						return nil
					}
					fmt.Println(rsp.Rate.String())
				} else {
					fmt.Println("github.com的连接异常：" + err.Error())
					return nil
				}
			}
			fmt.Println("gob源码从github.com中下载，github.com的连接正常")

			// 这里下面的client都是可用的了
			if rsp.LastPage != 0 {
				opts = &github.ListOptions{Page: rsp.LastPage, PerPage: perPage}
				releases, rsp, err = client.Repositories.ListReleases(context.Background(), Owner, Repo, opts)
				if err != nil {
					fmt.Println("任务终止：" + err.Error())
					return nil
				}
				fmt.Println(rsp.Rate.String())
			}
			fmt.Printf("最新的%v个版本\n", len(releases))
			for _, releaseTmp := range releases {
				fmt.Println(releaseTmp.GetTagName())
			}

			prompt := &survey.Input{
				Message: "请输入一个版本(更多可以参考 " + GitHubReleasesUrl + "，默认为最新版本)：",
			}
			if err = survey.AskOne(prompt, &version); err != nil {
				fmt.Println("任务终止：" + err.Error())
				return err
			}
			if version != "" {
				// 确认版本是否正确
				release, _, err = client.Repositories.GetReleaseByTag(context.Background(), Owner, Repo, version)
				if err != nil || release == nil {
					fmt.Println("版本不存在，创建应用失败，请参考 " + GitHubReleasesUrl)
					return nil
				}
			}
			if version == "" {
				release, _, err = client.Repositories.GetLatestRelease(context.Background(), Owner, Repo)
				if err != nil {
					fmt.Println("获取最新版本失败 " + err.Error())
					return nil
				}
				version = release.GetTagName()
			}
		}
		fmt.Println("====================================================")
		fmt.Println("开始进行创建应用操作")
		fmt.Println("创建目录：", folder)
		fmt.Println("应用名称：", mod)
		fmt.Println("gob框架版本：", release.GetTagName())

		templateFolder := filepath.Join(currentPath, "template-gob-"+version+"-"+cast.ToString(time.Now().Unix()))
		err := os.Mkdir(templateFolder, os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Println("创建临时目录", templateFolder)

		defer func() {
			if err := os.RemoveAll(templateFolder); err != nil {
				fmt.Println("删除临时文件夹错误：", err.Error())
			}
			fmt.Println("删除临时文件夹", templateFolder)
		}()

		// 拷贝template项目
		url := release.GetZipballURL()
		err = util.DownloadFile(filepath.Join(templateFolder, "template.zip"), url)
		if err != nil {
			return err
		}
		fmt.Println("下载zip包到template.zip")

		_, err = util.Unzip(filepath.Join(templateFolder, "template.zip"), templateFolder)
		if err != nil {
			return err
		}

		// 获取folder下的gob-xxx相关解压目录
		fInfos, err := os.ReadDir(templateFolder)
		if err != nil {
			return err
		}
		for _, fInfo := range fInfos {
			// 找到解压后的文件夹
			if fInfo.IsDir() && strings.Contains(fInfo.Name(), "gob-") {
				if err := os.Rename(filepath.Join(templateFolder, fInfo.Name()), folder); err != nil {
					return err
				}
			}
		}
		fmt.Println("解压zip包")

		_ = os.RemoveAll(path.Join(folder, ".git"))
		fmt.Println("删除.git目录")

		// 删除framework 目录
		_ = os.RemoveAll(path.Join(folder, "framework"))
		fmt.Println("删除framework目录")

		filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if path == filepath.Join(folder, "go.mod") {
				fmt.Println("更新文件:" + path)
				b = bytes.ReplaceAll(b, []byte("module github.com/chenbihao/gob"), []byte("module "+mod))
				b = bytes.ReplaceAll(b, []byte("require ("), []byte("require (\n\tgithub.com/chenbihao/gob "+version))
				err = os.WriteFile(path, b, 0644)
				if err != nil {
					return err
				}
				return nil
			}
			isContain := bytes.Contains(b, []byte("github.com/chenbihao/gob/app"))
			if isContain {
				fmt.Println("更新文件:" + path)
				b = bytes.ReplaceAll(b, []byte("github.com/chenbihao/gob/app"), []byte(mod+"/app"))
				err = os.WriteFile(path, b, 0644)
				if err != nil {
					return err
				}
			}
			return nil
		})
		fmt.Println("创建应用结束")
		fmt.Println("目录：", folder)
		fmt.Println("====================================================")
		return nil
	},
}
