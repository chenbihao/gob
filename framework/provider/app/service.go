package app

// 实现具体的服务实例 service.go

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/util"
	"github.com/google/uuid"
)

// AppService 代表 gob 框架的 App 实现
type AppService struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
	appID      string              // 表示当前这个app的唯一id, 可以用于分布式锁等
	// toolMode   bool                // 工具运行模式（通过 go install 安装至 $GOPATH/bin ）
	argsMap   map[string]string // 参数加载(--key=value)
	sysEnvMap map[string]string // 环境变量加载
	configMap map[string]string // 配置加载
}

var _ contract.App = (*AppService)(nil)

// NewGobApp 初始化 AppService
func NewGobApp(params ...any) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}
	// 有两个参数，一个是容器，一个是 baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)

	uid, _ := uuid.NewV7()
	appID := uid.String()
	configMap := map[string]string{}

	// toolMode := false
	// // 纯工具模式 ( 兼容 go install )
	// if os.Getenv("runMode") == "tool" || util.CheckBinaryFileInTheGOPATH() {
	// 	toolMode = true
	// }

	// gobApp := &AppService{baseFolder: baseFolder, container: container, appID: appID, configMap: configMap, toolMode: toolMode}
	gobApp := &AppService{baseFolder: baseFolder, container: container, appID: appID, configMap: configMap}
	_ = gobApp.loadEnvMaps()
	_ = gobApp.loadArgsMaps()
	return gobApp, nil

}

// AppID 表示当前这个app的唯一id, 可以用于分布式锁等
func (app *AppService) AppID() string {
	return app.appID
}

// Version 实现版本
func (app *AppService) Version() string {
	return framework.Version
}

// // IsToolMode 是否纯工具运行模式
// func (app *AppService) IsToolMode() bool {
// 	return app.toolMode
// }

// ---------------- 目录

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (app *AppService) BaseFolder() string {
	if app.baseFolder != "" {
		return app.baseFolder
	}
	baseFolder := app.getConfigBySequence("base_folder", "BASE_FOLDER", "app.path.base_folder")
	if baseFolder != "" {
		return baseFolder
	}

	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

// ---------------- 根目录下

// AppFolder 定义业务代码所在的目录，用于监控文件变更使用
func (app *AppService) AppFolder() string {
	val := app.getConfigBySequence("app_folder", "APP_FOLDER", "app.path.app_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.BaseFolder(), "app")
}

// ConfigFolder  表示配置文件地址
func (app *AppService) ConfigFolder() string {
	val := app.getConfigBySequence("config_folder", "CONFIG_FOLDER", "app.path.config_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.BaseFolder(), "config")
}

// StorageFolder 存储文件地址
func (app *AppService) StorageFolder() string {
	val := app.getConfigBySequence("storage_folder", "STORAGE_FOLDER", "app.path.storage_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.BaseFolder(), "storage")
}

// DeployFolder 存放部署的时候创建的文件夹
func (app AppService) DeployFolder() string {
	val := app.getConfigBySequence("deploy_folder", "DEPLOY_FOLDER", "app.path.deploy_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.BaseFolder(), "deploy")
}

// // TestFolder 定义测试需要的信息
// func (s *AppService) TestFolder() string {
// 	val := s.getConfigBySequence("test_folder", "TEST_FOLDER", "app.path.test_folder")
// 	if val != "" {
// 		return val
// 	}
// 	return filepath.Join(s.BaseFolder(), "test")
// }

// ---------------- app 目录下

// ProviderFolder 定义业务自己的通用服务提供者地址
func (app *AppService) ProviderFolder() string {
	val := app.getConfigBySequence("provider_folder", "PROVIDER_FOLDER", "app.path.provider_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.AppFolder(), "provider")
}

// // CommandFolder 定义业务自己的命令行服务提供者地址
// func (s *AppService) CommandFolder() string {
// 	val := s.getConfigBySequence("command_folder", "COMMAND_FOLDER", "app.path.command_folder")
// 	if val != "" {
// 		return val
// 	}
// 	return filepath.Join(s.AppFolder(), "command")
// }

// // HttpFolderr 定义业务自己的web服务提供者地址
// func (s *AppService) HttpFolder() string {
// 	val := s.getConfigBySequence("http_folder", "HTTP_FOLDER", "app.path.http_folder")
// 	if val != "" {
// 		return val
// 	}
// 	return filepath.Join(s.AppFolder(), "http")
// }

// // MiddlewareFolder 定义业务自己定义的中间件
// func (s *AppService) MiddlewareFolder() string {
// 	val := s.getConfigBySequence("middleware_folder", "MIDDLEWARE_FOLDER", "app.path.middleware_folder")
// 	if val != "" {
// 		return val
// 	}
// 	return filepath.Join(s.HttpFolder(), "middleware")
// }

// ---------------- storage 目录下

func (app *AppService) LogFolder() string {
	val := app.getConfigBySequence("log_folder", "LOG_FOLDER", "app.path.log_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.StorageFolder(), "log")
}

// RuntimeFolder 定义业务的运行中间态信息
func (app *AppService) RuntimeFolder() string {
	val := app.getConfigBySequence("runtime_folder", "RUNTIME_FOLDER", "app.path.runtime_folder")
	if val != "" {
		return val
	}
	return filepath.Join(app.StorageFolder(), "runtime")
}

// GetConfigByDefault 默认获取配置项的方法
// 配置优先级：参数>环境变量>配置文件
func (app *AppService) getConfigBySequence(argsKey string, sysEnvKey string, configKey string) string {
	if app.argsMap != nil {
		if val, ok := app.argsMap[argsKey]; ok {
			return val
		}
	}
	if app.sysEnvMap != nil {
		if val, ok := app.sysEnvMap[sysEnvKey]; ok {
			return val
		}
	}
	if app.configMap != nil {
		if val, ok := app.configMap[configKey]; ok {
			return val
		}
	}
	return ""
}

func (app *AppService) loadEnvMaps() error {
	if app.sysEnvMap == nil {
		app.sysEnvMap = map[string]string{}
	}
	// 读取环境变量
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		app.sysEnvMap[pair[0]] = pair[1]
	}
	return nil
}

func (app *AppService) loadArgsMaps() error {
	if app.argsMap == nil {
		app.argsMap = map[string]string{}
	}
	// load args, must format : --key=value
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--") {
			pair := strings.SplitN(strings.TrimPrefix(arg, "--"), "=", 2)
			if len(pair) == 2 {
				app.argsMap[pair[0]] = pair[1]
			}
		}
	}
	return nil
}

// LoadAppConfig 加载配置map
func (app *AppService) LoadAppConfig(kv map[string]string) {
	for key, val := range kv {
		app.configMap[key] = val
	}
}
