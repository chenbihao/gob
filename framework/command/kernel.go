package command

import (
	"github.com/chenbihao/gob/app/console"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/cobra"
	"github.com/chenbihao/gob/framework/command/model"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/robfig/cron/v3"
)

// RunRootCommand  初始化根Command并运行
func RunRootCommand(container framework.Container) error {
	// 根Command
	var rootCmd = &cobra.Command{
		// 定义根命令的关键字
		Use: "gob",
		// 简短介绍
		Short: "gob 命令",
		// 根命令的详细介绍
		Long: "gob 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令",
		// 根命令的执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现cobra默认的completion子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	// 为根Command设置服务容器
	rootCmd.SetContainer(container)

	// 绑定框架的命令
	if container.MustMake(contract.AppKey).(contract.App).IsToolMode() {
		AddKernelToolCommands(rootCmd)
	} else {
		AddKernelCommands(rootCmd)
		// 绑定业务的命令
		console.AddAppCommand(rootCmd)
	}

	// 执行RootCommand
	return rootCmd.Execute()
}

// AddKernelToolCommands will add all command/* to root command
func AddKernelToolCommands(root *cobra.Command) {
	root.AddCommand(initNewCommand())     // 挂载 new 命令
	root.AddCommand(initVersionCommand()) // 挂载 version 命令
}

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	InitCronCommands(root)
	root.AddCommand(initAppCommand())         // 挂载 app 命令
	root.AddCommand(initEnvCommand())         // 挂载 env 命令
	root.AddCommand(initCronCommand())        // 挂载 cron 命令
	root.AddCommand(initConfigCommand())      // 挂载 config 命令
	root.AddCommand(initBuildCommand())       // 挂载 build 命令
	root.AddCommand(initGoCommand())          // 挂载 go 命令
	root.AddCommand(initNpmCommand())         // 挂载 npm 命令
	root.AddCommand(initDevCommand())         // 挂载 dev 调试命令
	root.AddCommand(initProviderCommand())    // 挂载 provider 命令
	root.AddCommand(initCmdCommand())         // 挂载 command 命令
	root.AddCommand(initMiddlewareCommand())  // 挂载 middleware 命令
	root.AddCommand(initNewCommand())         // 挂载 new 命令
	root.AddCommand(initSwaggerCommand())     // 挂载 swagger 命令
	root.AddCommand(initDeployCommand())      // 挂载 deploy 命令
	root.AddCommand(initVersionCommand())     // 挂载 version 命令
	root.AddCommand(model.InitModelCommand()) // 挂载 model 命令
}

// InitCronCommands 初始化Cron相关的命令
func InitCronCommands(root *cobra.Command) {
	// 初始化cron相关命令
	if root.Cron == nil {
		// 初始化cron
		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
		root.CronSpecs = []cobra.CronSpec{}
	}
}
