package config

// 实现具体的服务实例 service.go

import (
	"bytes"
	"fmt"
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
	kstructs "github.com/knadh/koanf/providers/structs"
)

// ConfigService 是 Config 的具体实现
type ConfigService struct {
	c          framework.Container       // 容器
	lock       sync.RWMutex              // 配置文件读写锁
	keyDelim   string                    // key 路径的分隔符，默认为点
	folder     string                    // 文件夹
	kEnv       *koanf.Koanf              // 所有的环境变量
	kEnvStruct *contract.ConfigEnvStruct // gob 环境变量结构体
	kConfig    *koanf.Koanf              // 所有的配置
	kSubConfig map[string]*koanf.Koanf   // 所有的契约配置
}

var _ contract.Config = (*ConfigService)(nil)

func (c *ConfigService) GetEnv() *koanf.Koanf {
	return c.kEnv
}

func (c *ConfigService) GetEnvStruct() *contract.ConfigEnvStruct {
	return c.kEnvStruct
}

func (c *ConfigService) GetConfig() *koanf.Koanf {
	return c.kConfig
}

var delim = "."
var tagName = "conf"

// NewConfigService 初始化Config方法
func NewConfigService(params ...any) (any, error) {

	container := params[0].(framework.Container)
	appService := container.MustMake(contract.AppKey).(contract.App)

	// 读取环境变量（用于替换值）
	var kEnv = koanf.New(delim)
	// 读默认值
	_ = kEnv.Load(kstructs.Provider(contract.DefaultConfigEnvStruct, tagName), nil)
	// 读.env文件
	_ = kEnv.Load(kfile.Provider(filepath.Join(appService.BaseFolder(), ".env")), kdotenv.Parser())
	// 读环境变量
	var kSysEnv = koanf.New(delim)
	_ = kSysEnv.Load(kenv.Provider(delim, kenv.Opt{}), nil)
	_ = kEnv.Merge(kSysEnv)

	var envConfig = contract.ConfigEnvStruct{}
	_ = kEnv.UnmarshalWithConf("", &envConfig, koanf.UnmarshalConf{Tag: tagName})

	// 默认是极简模式，可选开启配置文件夹，可选开启部署配置分离模式（deploy_env：env/test/prod）
	configFolder := appService.BaseFolder()
	if envConfig.ConfigMode != "" {
		switch envConfig.ConfigMode {
		case contract.ConfigModeRoot:
			configFolder = filepath.Join(appService.BaseFolder(), envConfig.ConfigFolder) // 这个是全部配置都在一个文件中
		case contract.ConfigModeFolder:
			configFolder = filepath.Join(appService.BaseFolder(), envConfig.ConfigFolder) // 这个是可配置独立的配置文件
		case contract.ConfigModeDeploy:
			configFolder = filepath.Join(appService.BaseFolder(), envConfig.ConfigFolder, string(envConfig.AppEnv))
		}
	}
	fmt.Println("configFolder:", configFolder)

	// 初始化 config.yaml
	var k = koanf.New(delim)
	configFileName := "config"
	k = Load(configFolder, configFileName, k, kSysEnv)

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
		c:          container,
		lock:       sync.RWMutex{},
		keyDelim:   delim,
		folder:     configFolder,
		kEnv:       kEnv,
		kEnvStruct: &envConfig,
		kConfig:    k,
		//kSub:     make(map[string]*koanf.Koanf),
	}

	// todo 打印输出已选配置

	return gobConf, nil
}

func Load(configFolder string, configFileName string, k *koanf.Koanf, kSysEnv *koanf.Koanf) *koanf.Koanf {
	kConfig := kfile.Provider(filepath.Join(configFolder, configFileName+".yaml"))
	replaceAndLoad(k, kConfig, kSysEnv)
	// 监控文件夹文件
	_ = kConfig.Watch(func(event interface{}, err error) {
		if err != nil {
			log.Printf("watch error: %v", err)
			return
		}
		// Throw away the old config and load a fresh copy.
		log.Println("config changed. Reloading ...")
		k = koanf.New(delim)
		replaceAndLoad(k, kConfig, kSysEnv)
		k.Print()
	})
	// To stop a file watcher, call:
	// f.Unwatch()
	return k
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
