package command

import (
	"fmt"
	"github.com/chenbihao/gob/framework/cobra"
	"log"
	"os"
	"os/exec"
)

/*
## 命令介绍：
挂载 npm 命令
## 前置需求：无
## 支持命令：
```sh
./gob npm
```
## 支持配置：无
*/

// 用于生成文档定位说明
const NpmGoCommandKey = "npm 命令"

func initNpmCommand() *cobra.Command {
	return npmCommand
}

// npm just run local go bin
var npmCommand = &cobra.Command{
	Use:   "npm",
	Short: "运行 PATH/npm 的命令",
	RunE: func(c *cobra.Command, args []string) error {
		fmt.Println("=============  执行 npm 命令 ============")
		path, err := exec.LookPath("npm")
		if err != nil {
			log.Fatalln("gob npm: should install npm in your PATH")
		}

		cmd := exec.Command(path, args...)
		cmd.Dir = frontendFolder
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return nil
	},
}
