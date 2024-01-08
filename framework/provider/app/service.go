package app

// 实现具体的服务实例 service.go

import (
	"errors"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/util"
	"github.com/google/uuid"
	flag "github.com/spf13/pflag"
	"path/filepath"
)

// AppService 代表 gob 框架的 App 实现
type AppService struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
	appID      string              // 表示当前这个app的唯一id, 可以用于分布式锁等

	configMap map[string]string // 配置加载
}

var _ contract.App = (*AppService)(nil)

// NewGobApp 初始化 AppService
func NewGobApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}
	// 有两个参数，一个是容器，一个是 baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)

	// 如果没有设置，则使用参数
	if baseFolder == "" {
		flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数, 默认为当前路径")
		flag.Parse()
	}

	appID := uuid.New().String()
	return &AppService{baseFolder: baseFolder, container: container, appID: appID}, nil
}

// AppID 表示当前这个app的唯一id, 可以用于分布式锁等
func (s *AppService) AppID() string {
	return s.appID
}

// Version 实现版本
func (s *AppService) Version() string {
	return "0.1.6"
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (s *AppService) BaseFolder() string {
	if s.baseFolder != "" {
		return s.baseFolder
	}

	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

// ConfigFolder  表示配置文件地址
func (s *AppService) ConfigFolder() string {
	if val, ok := s.configMap["config_folder"]; ok {
		return val
	}
	return filepath.Join(s.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (s *AppService) LogFolder() string {
	if val, ok := s.configMap["log_folder"]; ok {
		return val
	}
	return filepath.Join(s.StorageFolder(), "log")
}

func (s *AppService) HttpFolder() string {
	if val, ok := s.configMap["http_folder"]; ok {
		return val
	}
	return filepath.Join(s.BaseFolder(), "http")
}

func (s *AppService) ConsoleFolder() string {
	if val, ok := s.configMap["console_folder"]; ok {
		return val
	}
	return filepath.Join(s.BaseFolder(), "console")
}

func (s *AppService) StorageFolder() string {
	if val, ok := s.configMap["storage_folder"]; ok {
		return val
	}
	return filepath.Join(s.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (s *AppService) ProviderFolder() string {
	if val, ok := s.configMap["provider_folder"]; ok {
		return val
	}
	return filepath.Join(s.BaseFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (s *AppService) MiddlewareFolder() string {
	if val, ok := s.configMap["middleware_folder"]; ok {
		return val
	}
	return filepath.Join(s.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (s *AppService) CommandFolder() string {
	if val, ok := s.configMap["command_folder"]; ok {
		return val
	}
	return filepath.Join(s.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (s *AppService) RuntimeFolder() string {
	if val, ok := s.configMap["runtime_folder"]; ok {
		return val
	}
	return filepath.Join(s.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (s *AppService) TestFolder() string {
	if val, ok := s.configMap["test_folder"]; ok {
		return val
	}
	return filepath.Join(s.BaseFolder(), "test")
}

// LoadAppConfig 加载配置map
func (s *AppService) LoadAppConfig(kv map[string]string) {
	for key, val := range kv {
		s.configMap[key] = val
	}
}
