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
	root.AddCommand(initGoCommand())
	// 挂载 npm 命令
	root.AddCommand(initNpmCommand())
	// 挂载 dev 调试命令
	root.AddCommand(initDevCommand())
	// 挂载 provider 命令
	root.AddCommand(initProviderCommand())
	// 挂载 command 命令
	root.AddCommand(initCmdCommand())
	// 挂载 middleware 命令
	root.AddCommand(initMiddlewareCommand())
	// 挂载 new 命令
	root.AddCommand(initNewCommand())
	// 挂载 swagger 命令
	root.AddCommand(initSwaggerCommand())
}

// AddNewCommands will add new command to root command
func AddNewCommands(root *cobra.Command) {
	// 挂载 new 命令
	root.AddCommand(initNewCommand())
}
