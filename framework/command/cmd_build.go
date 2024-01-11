package command

import (
	"fmt"
	"github.com/chenbihao/gob/framework/cobra"
	"log"
	"os/exec"
)

// 前端文件夹
const frontendFolder = "./gob_frontend/"

// build相关的命令
func initBuildCommand() *cobra.Command {
	buildCommand.AddCommand(buildSelfCommand)
	buildCommand.AddCommand(buildBackendCommand)
	buildCommand.AddCommand(buildFrontendCommand)
	buildCommand.AddCommand(buildAllCommand)
	return buildCommand
}

var buildCommand = &cobra.Command{
	Use:   "build",
	Short: "编译相关命令",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			c.Help()
		}
		return nil
	},
}

var buildSelfCommand = &cobra.Command{
	Use:   "self",
	Short: "编译 gob 命令",
	RunE: func(c *cobra.Command, args []string) error {
		fmt.Println("=============  后端编译开始 ============")
		path, err := exec.LookPath("go")
		if err != nil {
			log.Fatalln("请安装 go 在你的 PATH 路径下")
		}

		cmd := exec.Command(path, "build", "-o", "gob", "./")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			fmt.Println("=============  后端编译失败 ============")
			return err
		}
		fmt.Println("build success please run ./gob direct")
		fmt.Println("=============  后端编译成功 ============")
		return nil
	},
}

var buildBackendCommand = &cobra.Command{
	Use:   "backend",
	Short: "使用 go 编译后端",
	RunE: func(c *cobra.Command, args []string) error {
		return buildSelfCommand.RunE(c, args)
	},
}

var buildFrontendCommand = &cobra.Command{
	Use:   "frontend",
	Short: "使用 npm 编译前端",
	RunE: func(c *cobra.Command, args []string) error {
		fmt.Println("=============  前端编译开始 ============")
		// 获取path路径下的npm命令
		path, err := exec.LookPath("npm")
		if err != nil {
			log.Fatalln("请安装 npm 在你的 PATH 路径下")
		}

		// 执行npm run build
		cmd := exec.Command(path, "run", "build")
		cmd.Dir = frontendFolder
		// 将输出保存在out中
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			fmt.Println("=============  前端编译失败 ============")
			return err
		}
		// 打印输出
		fmt.Print(string(out))
		fmt.Println("=============  前端编译成功 ============")
		return nil
	},
}

var buildAllCommand = &cobra.Command{
	Use:   "all",
	Short: "同时编译前端和后端",
	RunE: func(c *cobra.Command, args []string) error {
		err := buildFrontendCommand.RunE(c, args)
		if err != nil {
			return err
		}
		err = buildBackendCommand.RunE(c, args)
		if err != nil {
			return err
		}
		return nil
	},
}
