package command

import (
	"fmt"
	"github.com/chenbihao/gob/framework/cobra"
	"log"
	"os"
	"os/exec"
)

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
