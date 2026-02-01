package contract

import (
	"github.com/knadh/koanf/v2"
	"path/filepath"
)

/*
## 服务介绍：
提供基础的配置文件获取方法
## 支持命令：
[config](../command/config)
## 支持配置：
[config](../config/config)
*/

// ConfigKey 是配置服务字符串凭证
const ConfigKey = "gob:config"

// Config 定义了配置文件服务，读取配置文件，支持点分割的路径读取
// 例如: .Get("app.name") 表示从 app 文件中读取 name 属性
// 建议使用 yaml 属性, https://yaml.org/spec/1.2/spec.html
type Config interface {
	GetEnv() *koanf.Koanf
	GetEnvStruct() *ConfigEnvStruct
	GetConfig() *koanf.Koanf
	GetSubConfig(key string) *koanf.Koanf
}

type ConfigEnvStruct struct {
	ConfigMode ConfigMode `conf:"configMode"` // 默认 "root"，可选 configMode
	AppEnv     AppEnv     `conf:"appEnv"`     // 默认 "dev"，当 configMode = Deploy 时生效

	ConfigFolder     string `conf:"configFolder"`     // configFolder 定义配置文件所在的目录
	AppFolder        string `conf:"appFolder"`        // appFolder 定义业务代码所在的目录，用于监控文件变更使用
	TestFolder       string `conf:"testFolder"`       // testFolder 定义测试需要的信息
	StorageFolder    string `conf:"storageFolder"`    // storageFolder 存储文件地址
	DeployFolder     string `conf:"deployFolder"`     // deployFolder 存放部署的时候创建的文件夹
	ProviderFolder   string `conf:"providerFolder"`   // providerFolder 定义业务自己的通用服务提供者地址
	HttpFolder       string `conf:"httpFolder"`       // httpFolder 定义业务自己的web服务提供者地址
	MiddlewareFolder string `conf:"middlewareFolder"` // middlewareFolder 定义业务自己定义的中间件
	CommandFolder    string `conf:"commandFolder"`    // commandFolder 定义业务自己的命令行服务提供者地址
	WailsFolder      string `conf:"WailsFolder"`      // wailsFolder 定义业务自己的app服务提供者地址
	LogFolder        string `conf:"logFolder"`        // logFolder
	RuntimeFolder    string `conf:"runtimeFolder"`    // runtimeFolder 定义业务的运行中间态信息
}

type ConfigMode string

const (
	ConfigModeRoot   ConfigMode = "root"   // root 读取根目录的 config.yaml
	ConfigModeFolder            = "folder" // folder 读取 {config} 目录下读取，支持配置分离
	ConfigModeDeploy            = "deploy" // deploy 读取 {config}/{} 目录下读取，支持配置分离（对应的部署环境配置目录如dev/test/prod）
)

type AppEnv string

const (
	AppEnvDev  AppEnv = "dev"  // dev 代表开发环境
	AppEnvTest        = "test" // test 代表测试环境
	AppEnvProd        = "prod" // prod 代表生产环境
)

type ConfigFolder struct {
}

var DefaultConfigEnvStruct = ConfigEnvStruct{
	ConfigMode:       ConfigModeDeploy,
	AppEnv:           AppEnvDev,
	ConfigFolder:     filepath.Join("config"),
	AppFolder:        filepath.Join("app"),
	TestFolder:       filepath.Join("test"),
	StorageFolder:    filepath.Join("storage"),
	DeployFolder:     filepath.Join("deploy"),
	ProviderFolder:   filepath.Join("app", "provider"),
	HttpFolder:       filepath.Join("app", "http"),
	MiddlewareFolder: filepath.Join("app", "http", "middleware"),
	CommandFolder:    filepath.Join("app", "command"),
	WailsFolder:      filepath.Join("app", "wails"),
	LogFolder:        filepath.Join("storage", "log"),
	RuntimeFolder:    filepath.Join("storage", "runtime"),
}
