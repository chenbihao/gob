package command

import (
	"fmt"
	"github.com/chenbihao/gob/framework/cobra"
	"log"
	"os"
	"os/exec"
)

func initGoCommand() *cobra.Command {
	return goCommand
}

// go just run local go bin
var goCommand = &cobra.Command{
	Use:   "go",
	Short: "运行 path/go 程序，要求go 必须安装",
	RunE: func(c *cobra.Command, args []string) error {
		fmt.Println("=============  执行 go 命令 ============")
		path, err := exec.LookPath("go")
		if err != nil {
			log.Fatalln("gob go: should install go in your PATH")
		}

		cmd := exec.Command(path, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return nil
	},
}
