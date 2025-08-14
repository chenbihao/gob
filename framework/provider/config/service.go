package config

// 实现具体的服务实例 service.go

import (
	"bytes"
	"github.com/knadh/koanf/v2"
	"log"
	"path/filepath"
	"sync"

	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"

	kdotenv "github.com/knadh/koanf/parsers/dotenv"
	kyaml "github.com/knadh/koanf/parsers/yaml"

	kenv "github.com/knadh/koanf/providers/env/v2"
	kfile "github.com/knadh/koanf/providers/file"
	krawbytes "github.com/knadh/koanf/providers/rawbytes"
)

// ConfigService 是 Config 的具体实现
type ConfigService struct {
	c          framework.Container     // 容器
	lock       sync.RWMutex            // 配置文件读写锁
	keyDelim   string                  // key 路径的分隔符，默认为点
	folder     string                  // 文件夹
	kEnv       *koanf.Koanf            // 所有的环境变量
	kConfig    *koanf.Koanf            // 所有的配置
	kSubConfig map[string]*koanf.Koanf // 所有的契约配置
}

var _ contract.Config = (*ConfigService)(nil)

type ConfigMode string // 定义自定义类型作为枚举基础
const (
	Root      ConfigMode = "root"      // 读取根目录的 config.yaml
	Folder               = "folder"    // 读取 config 目录下读取，支持配置分离
	DeployEnv            = "deployEnv" // 读取 config/{deploy_env} 目录下读取，支持配置分离（对应的部署环境配置目录如dev/test/prod）
)

const (
	EnvConfigMode = "configMode"
	EnvDeployEnv  = "deployEnv"
)

// NewConfigService 初始化Config方法
func NewConfigService(params ...any) (any, error) {

	container := params[0].(framework.Container)
	appService := container.MustMake(contract.AppKey).(contract.App)

	delim := "."

	// 读取环境变量（先读取环境变量，后读取.env文件替换值）
	var kEnv = koanf.New(delim)
	_ = kEnv.Load(kenv.Provider(delim, kenv.Opt{}), nil)

	var kEnvFile = koanf.New(delim)
	_ = kEnvFile.Load(kfile.Provider(filepath.Join(appService.BaseFolder(), ".env")), kdotenv.Parser())
	_ = kEnv.Merge(kEnvFile)

	// 默认是极简模式，可选开启配置文件夹，可选开启部署配置分离模式（deploy_env：env/test/prod）
	configMode := kEnv.Get(EnvConfigMode)
	configFolder := appService.BaseFolder()
	if configMode != nil {
		switch configMode {
		case Root:
			configFolder = appService.BaseFolder()
		case Folder:
			configFolder = filepath.Join(appService.BaseFolder(), "config")
		case DeployEnv:
			deployEnv := kEnv.String(EnvDeployEnv)
			if deployEnv == "" {
				deployEnv = "dev"
			}
			configFolder = filepath.Join(appService.BaseFolder(), "config", deployEnv)
		}
	}

	// 初始化 config.yaml
	var k = koanf.New(delim)
	kConfig := kfile.Provider(filepath.Join(configFolder, "config.yaml"))
	replaceAndLoad(k, kConfig, kEnv)
	// 监控文件夹文件
	_ = kConfig.Watch(func(event interface{}, err error) {
		if err != nil {
			log.Printf("watch error: %v", err)
			return
		}
		// Throw away the old config and load a fresh copy.
		log.Println("config changed. Reloading ...")
		k = koanf.New(delim)
		replaceAndLoad(k, kConfig, kEnv)
		k.Print()
	})
	// To stop a file watcher, call:
	// f.Unwatch()

	//var kApp = koanf.New(delim)
	//kAppFile := kfile.Provider(appService.BaseFolder() + "/config/dev/app.yaml")
	//kAppByte, _ := kAppFile.ReadBytes()
	//_ = kApp.Load(krawbytes.Provider(kAppByte), kyaml.Parser())
	//_ = k.MergeAt(kApp, "app")
	//
	//var kCache = koanf.New(delim)
	//_ = kCache.Load(kfile.Provider(appService.BaseFolder()+"/config/dev/cache.yaml"), kyaml.Parser())
	//_ = k.MergeAt(kCache, "cache")
	//k.Print()

	// 实例化
	gobConf := &ConfigService{
		c:        container,
		lock:     sync.RWMutex{},
		keyDelim: delim,
		folder:   configFolder,
		kEnv:     kEnv,
		kConfig:  k,
		//kSub:     make(map[string]*koanf.Koanf),
	}

	return gobConf, nil
}

// replaceAndLoad 替换环境变量maps并加载配置
func replaceAndLoad(k *koanf.Koanf, kConfig *kfile.File, kEnv *koanf.Koanf) []byte {
	// todo Such scenarios will need mutex locking.
	kConfigByte, _ := kConfig.ReadBytes()
	kConfigByte = replaceEnvKey(kConfigByte, kEnv.StringMap(""))
	_ = k.Load(krawbytes.Provider(kConfigByte), kyaml.Parser())
	return kConfigByte
}

// replaceEnvKey 表示使用环境变量maps替换context中的env(xxx)的环境变量
func replaceEnvKey(content []byte, maps map[string]string) []byte {
	if maps == nil {
		return content
	}
	// 直接使用ReplaceAll替换。这个性能可能不是最优，但是配置文件加载，频率是比较低的，可以接受
	for key, val := range maps {
		reKey := "env(" + key + ")"
		content = bytes.ReplaceAll(content, []byte(reKey), []byte(val))
	}
	return content
}
