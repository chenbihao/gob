from https://github.com/spf13/cobra/tree/v1.9.1

gob adaptation :

- add `gob_*.go`
- replace `"github.com/spf13/cobra"` -> `"github.com/chenbihao/gob/framework/cobra"`

- `command.go` edit:

```go
type Command struct {

	// gob改动：引入定时库
	Cron      *cron.Cron // Command支持cron，只在RootCommand中有这个值
	CronSpecs []CronSpec // 对应Cron命令的信息

	// gob改动：引入服务容器
	container framework.Container

```









