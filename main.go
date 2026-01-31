package main

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/cobra"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/provider/app"
	"github.com/chenbihao/gob/framework/provider/config"
	"github.com/chenbihao/gob/framework/provider/id"
)

func main() {

	// 初始化服务容器
	container := framework.NewGobContainer()

	_ = container.Bind(&app.AppProvider{})
	_ = container.Bind(&config.ConfigProvider{})
	_ = container.Bind(&id.IDProvider{})

	appService := container.MustMake(contract.AppKey).(contract.App)
	_ = container.MustMake(contract.ConfigKey).(contract.Config)

	println("AppID：" + appService.AppID())
	println("BaseFolder：" + appService.BaseFolder())

	<-make(chan struct{})

	//// 将 HTTP 引擎初始化,并且作为服务提供者绑定到服务容器中
	//if engine, err := http.NewHttpEngine(container); err == nil {
	//	// 绑定 Kernel 服务提供者
	//	_ = container.Bind(&kernel.KernelProvider{HttpEngine: engine})
	//}

	// 运行root命令
	//_ = RunRootCommand(container)
}

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

	// // 绑定框架的命令
	// command.AddKernelCommands(rootCmd)
	// // 绑定业务的命令
	// AddAppCommand(rootCmd)

	// 执行RootCommand
	return rootCmd.Execute()
}
