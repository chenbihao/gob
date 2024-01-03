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

// GobAppService 代表 gob 框架的 App 实现
type GobAppService struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
}

var _ contract.App = (*GobAppService)(nil)

// Version 实现版本
func (s GobAppService) Version() string {
	return "0.1.2"
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (s GobAppService) BaseFolder() string {
	if s.baseFolder != "" {
		return s.baseFolder
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
func (s GobAppService) ConfigFolder() string {
	return filepath.Join(s.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (s GobAppService) LogFolder() string {
	return filepath.Join(s.StorageFolder(), "log")
}

func (s GobAppService) HttpFolder() string {
	return filepath.Join(s.BaseFolder(), "http")
}

func (s GobAppService) ConsoleFolder() string {
	return filepath.Join(s.BaseFolder(), "console")
}

func (s GobAppService) StorageFolder() string {
	return filepath.Join(s.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (s GobAppService) ProviderFolder() string {
	return filepath.Join(s.BaseFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (s GobAppService) MiddlewareFolder() string {
	return filepath.Join(s.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (s GobAppService) CommandFolder() string {
	return filepath.Join(s.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (s GobAppService) RuntimeFolder() string {
	return filepath.Join(s.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (s GobAppService) TestFolder() string {
	return filepath.Join(s.BaseFolder(), "test")
}

// NewGobApp 初始化 GobAppService
func NewGobApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}
	// 有两个参数，一个是容器，一个是 baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return GobAppService{baseFolder: baseFolder, container: container}, nil // todo 这里可能得规范下返回的是指针或者实体
}
