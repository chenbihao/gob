package app

// 实现具体的服务实例 service.go

import (
	"errors"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/util"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strings"
)

// AppService 代表 gob 框架的 App 实现
type AppService struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
	appID      string              // 表示当前这个app的唯一id, 可以用于分布式锁等

	configMap map[string]string // 配置加载
	envMap    map[string]string // 环境变量加载
	argsMap   map[string]string // 参数加载
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

	appID := uuid.New().String()
	configMap := map[string]string{}

	gobApp := &AppService{baseFolder: baseFolder, container: container, appID: appID, configMap: configMap}
	_ = gobApp.loadEnvMaps()
	_ = gobApp.loadArgsMaps()
	return gobApp, nil

}

// AppID 表示当前这个app的唯一id, 可以用于分布式锁等
func (s *AppService) AppID() string {
	return s.appID
}

// Version 实现版本
func (s *AppService) Version() string {
	return GobVersion
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (s *AppService) BaseFolder() string {
	if s.baseFolder != "" {
		return s.baseFolder
	}
	baseFolder := s.getConfigBySequence("base_folder", "BASE_FOLDER", "app.path.base_folder")
	if baseFolder != "" {
		return baseFolder
	}
	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

// ---------------- 根目录下

// AppFolder 定义业务代码所在的目录，用于监控文件变更使用
func (s *AppService) AppFolder() string {
	val := s.getConfigBySequence("app_folder", "APP_FOLDER", "app.path.app_folder")
	if val != "" {
		return val
	}
	return filepath.Join(s.BaseFolder(), "app")
}

// ConfigFolder  表示配置文件地址
func (s *AppService) ConfigFolder() string {
	val := s.getConfigBySequence("config_folder", "CONFIG_FOLDER", "app.path.config_folder")
	if val != "" {
		return val
	}
	return filepath.Join(s.BaseFolder(), "config")
}

// StorageFolder 存储文件地址
func (s *AppService) StorageFolder() string {
	val := s.getConfigBySequence("storage_folder", "STORAGE_FOLDER", "app.path.storage_folder")
	if val != "" {
		return val
	}
	return filepath.Join(s.BaseFolder(), "storage")
}

// TestFolder 定义测试需要的信息
func (s *AppService) TestFolder() string {
	val := s.getConfigBySequence("test_folder", "TEST_FOLDER", "app.path.test_folder")
	if val != "" {
		return val
	}
	return filepath.Join(s.BaseFolder(), "test")
}

// DeployFolder 存放部署的时候创建的文件夹
func (s AppService) DeployFolder() string {
	val := s.getConfigBySequence("deploy_folder", "DEPLOY_FOLDER", "app.path.deploy_folder")
	if val != "" {
		return val
	}
	return filepath.Join(s.BaseFolder(), "deploy")
}

// ---------------- app 目录下

// ConsoleFolderr 定义业务自己的命令行服务提供者地址
func (s *AppService) ConsoleFolder() string {
	val := s.getConfigBySequence("console_folder", "CONSOLE_FOLDER", "app.path.console_folder")
	if val != "" {
		return val
	}
	return filepath.Join(s.AppFolder(), "console")
}

// HttpFolderr 定义业务自己的web服务提供者地址
func (s *AppService) HttpFolder() string {
	val := s.getConfigBySequence("http_folder", "HTTP_FOLDER", "app.path.http_folder")
	if val != "" {
		return val
	}
	return filepath.Join(s.AppFolder(), "http")
}

// ProviderFolder 定义业务自己的通用服务提供者地址
func (s *AppService) ProviderFolder() string {
	val := s.getConfigBySequence("provider_folder", "PROVIDER_FOLDER", "app.path.provider_folder")
	if val != "" {
		return val
	}
	return filepath.Join(s.AppFolder(), "provider")
}

// CommandFolder 定义业务定义的命令
func (s *AppService) CommandFolder() string {
	val := s.getConfigBySequence("command_folder", "COMMAND_FOLDER", "app.path.command_folder")
	if val != "" {
		return val
	}
	return filepath.Join(s.ConsoleFolder(), "command")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (s *AppService) MiddlewareFolder() string {
	val := s.getConfigBySequence("middleware_folder", "MIDDLEWARE_FOLDER", "app.path.middleware_folder")
	if val != "" {
		return val
	}
	return filepath.Join(s.HttpFolder(), "middleware")
}

// ---------------- config 目录下

// ---------------- storage 目录下

func (s *AppService) LogFolder() string {
	val := s.getConfigBySequence("log_folder", "LOG_FOLDER", "app.path.log_folder")
	if val != "" {
		return val
	}
	return filepath.Join(s.StorageFolder(), "log")
}

// RuntimeFolder 定义业务的运行中间态信息
func (s *AppService) RuntimeFolder() string {
	val := s.getConfigBySequence("runtime_folder", "RUNTIME_FOLDER", "app.path.runtime_folder")
	if val != "" {
		return val
	}
	return filepath.Join(s.StorageFolder(), "runtime")
}

// GetConfigByDefault 默认获取配置项的方法
// 配置优先级：参数>环境变量>配置文件
func (s *AppService) getConfigBySequence(argsKey string, envKey string, configKey string) string {
	if s.argsMap != nil {
		if val, ok := s.argsMap[argsKey]; ok {
			return val
		}
	}
	if s.envMap != nil {
		if val, ok := s.envMap[envKey]; ok {
			return val
		}
	}
	if s.configMap != nil {
		if val, ok := s.configMap[configKey]; ok {
			return val
		}
	}
	return ""
}

func (s *AppService) loadEnvMaps() error {
	if s.envMap == nil {
		s.envMap = map[string]string{}
	}
	// 读取环境变量
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		s.envMap[pair[0]] = pair[1]
	}
	return nil
}

func (s *AppService) loadArgsMaps() error {
	if s.argsMap == nil {
		s.argsMap = map[string]string{}
	}
	// load args, must format : --key=value
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--") {
			pair := strings.SplitN(strings.TrimPrefix(arg, "--"), "=", 2)
			if len(pair) == 2 {
				s.argsMap[pair[0]] = pair[1]
			}
		}
	}
	return nil
}

// LoadAppConfig 加载配置map
func (s *AppService) LoadAppConfig(kv map[string]string) {
	for key, val := range kv {
		s.configMap[key] = val
	}
}
