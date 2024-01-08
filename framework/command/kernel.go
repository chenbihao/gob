package command

import "github.com/chenbihao/gob/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	// 挂载 AppCommand 命令
	root.AddCommand(initAppCommand())

	// 挂载 cron 命令
	root.AddCommand(initCronCommand())

	// 挂载 环境变量 命令
	root.AddCommand(initEnvCommand())
}
