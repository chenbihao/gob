package command

import "github.com/chenbihao/gob/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	// 挂载 app 命令
	root.AddCommand(initAppCommand())
	// 挂载 cron 命令
	root.AddCommand(initCronCommand())
	// 挂载 env 命令
	root.AddCommand(initEnvCommand())
	// 挂载 config 命令
	root.AddCommand(initConfigCommand())
	// 挂载 build 命令
	root.AddCommand(initBuildCommand())
	// 挂载 go 命令
	root.AddCommand(goCommand)
	// 挂载 npm 命令
	root.AddCommand(npmCommand)
}
