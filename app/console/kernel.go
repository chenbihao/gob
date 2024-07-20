package console

import (
	"github.com/chenbihao/gob/app/console/command/demo"
	"github.com/chenbihao/gob/framework/cobra"
)

// 绑定业务的命令（业务定义的命令我们使用 `app/console/kernel.go` 中的 `AddAppCommand`进行挂载）
func AddAppCommand(rootCmd *cobra.Command) {
	// demo 例子
	rootCmd.AddCommand(demo.InitFoo())

	// 每秒调用一次Foo命令
	//rootCmd.AddCronCommand("* * * * * *", demo.FooCommand)

	// 启动一个分布式任务调度，调度的服务名称为init_func_for_test
	// 每个节点每5s调用一次Foo命令，抢占到了调度任务的节点将抢占锁持续挂载2s才释放
	//rootCmd.AddDistributedCronCommand("foo_func_for_test",
	//	"*/5 * * * * *", demo.FooCommand, 2*time.Second)
}
