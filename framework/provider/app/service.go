package app

// 实现具体的服务实例 service.go

import (
	"errors"
	"flag"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/util"
	"path/filepath"
)

// GobApp 代表 gob 框架的 App 实现
type GobApp struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
}

var _ contract.App = (*GobApp)(nil)

// Version 实现版本
func (app GobApp) Version() string {
	return "0.1.1"
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (app GobApp) BaseFolder() string {
	if app.baseFolder != "" {
		return app.baseFolder
	}
	// 如果没有设置，则使用参数
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder 参数, 默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}
	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

// ConfigFolder  表示配置文件地址
func (app GobApp) ConfigFolder() string {
	return filepath.Join(app.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (app GobApp) LogFolder() string {
	return filepath.Join(app.StorageFolder(), "log")
}

func (app GobApp) HttpFolder() string {
	return filepath.Join(app.BaseFolder(), "http")
}

func (app GobApp) ConsoleFolder() string {
	return filepath.Join(app.BaseFolder(), "console")
}

func (app GobApp) StorageFolder() string {
	return filepath.Join(app.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (app GobApp) ProviderFolder() string {
	return filepath.Join(app.BaseFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (app GobApp) MiddlewareFolder() string {
	return filepath.Join(app.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (app GobApp) CommandFolder() string {
	return filepath.Join(app.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (app GobApp) RuntimeFolder() string {
	return filepath.Join(app.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (app GobApp) TestFolder() string {
	return filepath.Join(app.BaseFolder(), "test")
}

// NewGobApp 初始化 GobApp
func NewGobApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}
	// 有两个参数，一个是容器，一个是 baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return GobApp{baseFolder: baseFolder, container: container}, nil // todo 这里可能得规范下返回的是指针或者实体
}
