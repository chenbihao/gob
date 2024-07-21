package command

import (
	"github.com/chenbihao/gob/framework/cobra"
	"github.com/chenbihao/gob/framework/command/model"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/robfig/cron/v3"
)

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	if root.GetContainer().MustMake(contract.AppKey).(contract.App).IsToolMode() {
		AddKernelOnlyToolCommands(root)
	} else {
		AddKernelAllCommands(root)
	}
}

// AddKernelToolCommands 添加 纯工具模式支持的命令到 root （go install 时运行的全局工具命令，尚未适配）
func AddKernelOnlyToolCommands(root *cobra.Command) {
	root.AddCommand(initNewCommand())     // 挂载 new 命令
	root.AddCommand(initVersionCommand()) // 挂载 version 命令
}

// AddKernelToolCommands 添加 command/* 到 root
func AddKernelAllCommands(root *cobra.Command) {
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
