package config

// 实现具体的服务实例 service.go

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/knadh/koanf/v2"

	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"

	kdotenv "github.com/knadh/koanf/parsers/dotenv"
	kyaml "github.com/knadh/koanf/parsers/yaml"
	kconfmap "github.com/knadh/koanf/providers/confmap"

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
var _ contract.SubConfigRegistry = (*ConfigService)(nil)

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
	if err := kEnv.Load(kstructs.Provider(contract.DefaultConfigEnvStruct, tagName), nil); err != nil {
		log.Printf("load default env values error: %v", err)
	}
	// 读.env文件
	envFilePath := filepath.Join(appService.BaseFolder(), ".env")
	if _, err := os.Stat(envFilePath); err == nil {
		if err := kEnv.Load(kfile.Provider(envFilePath), kdotenv.Parser()); err != nil {
			log.Printf("load .env file error: %v", err)
		}
	}
	// 读环境变量
	var kSysEnv = koanf.New(delim)
	if err := kSysEnv.Load(kenv.Provider(delim, kenv.Opt{}), nil); err != nil {
		log.Printf("load system env error: %v", err)
	}
	if err := kEnv.Merge(kSysEnv); err != nil {
		log.Printf("merge env error: %v", err)
	}

	var envConfig = contract.ConfigEnvStruct{}
	if err := kEnv.UnmarshalWithConf("", &envConfig, koanf.UnmarshalConf{Tag: tagName}); err != nil {
		log.Printf("unmarshal env config error: %v", err)
	}

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

	// 先创建 ConfigService 实例，kConfig 在 Load 中初始化
	configFileName := "config"
	gobConf := &ConfigService{
		c:          container,
		lock:       sync.RWMutex{},
		keyDelim:   delim,
		folder:     configFolder,
		kEnv:       kEnv,
		kEnvStruct: &envConfig,
		kSubConfig: make(map[string]*koanf.Koanf),
	}

	// 初始化 config.yaml，传入 ConfigService 以支持热重载时更新 kConfig
	Load(configFolder, configFileName, gobConf, kSysEnv)

	// 打印输出已选配置
	log.Printf("Config initialized - Mode: %s, Folder: %s", envConfig.ConfigMode, configFolder)

	return gobConf, nil
}

func Load(configFolder string, configFileName string, c *ConfigService, kSysEnv *koanf.Koanf) *koanf.Koanf {
	// 先初始化 kConfig，避免 nil 指针
	if c.kConfig == nil {
		c.kConfig = koanf.New(delim)
	}

	configFilePath := filepath.Join(configFolder, configFileName+".yaml")

	// 尝试从文件加载配置
	kConfig := kfile.Provider(configFilePath)
	replaceAndLoad(c.kConfig, kConfig, kSysEnv)

	// 监控文件夹文件
	_ = kConfig.Watch(func(event interface{}, err error) {
		if err != nil {
			log.Printf("watch error: %v", err)
			return
		}
		log.Println("config changed. Reloading ...")
		c.lock.Lock()
		defer c.lock.Unlock()
		// 使用新的 Koanf 实例替换旧的
		c.kConfig = koanf.New(delim)
		replaceAndLoad(c.kConfig, kConfig, kSysEnv)
		c.kConfig.Print()
	})

	return c.kConfig
}

// replaceAndLoad 替换环境变量maps并加载配置
func replaceAndLoad(k *koanf.Koanf, kConfig *kfile.File, kEnv *koanf.Koanf) []byte {
	kConfigByte, err := kConfig.ReadBytes()
	if err != nil {
		log.Printf("read config bytes error: %v", err)
		return nil
	}
	kConfigByte = replaceEnvKey(kConfigByte, kEnv.StringMap(""))
	if err := k.Load(krawbytes.Provider(kConfigByte), kyaml.Parser()); err != nil {
		log.Printf("load config error: %v", err)
		return nil
	}
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

// RegisterSubConfig 注册一个 ServiceConfig
// 优先级：环境变量 > 子配置文件 > 主配置文件 > 代码默认值
func (c *ConfigService) RegisterSubConfig(config framework.ServiceConfig) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	key := config.ConfigName()
	if _, exists := c.kSubConfig[key]; exists {
		return fmt.Errorf("sub config %s already registered", key)
	}

	// 创建独立的 Koanf 实例
	k := koanf.New(delim)

	// 1. 加载默认值（最低优先级）
	if defaults := config.Defaults(); len(defaults) > 0 {
		if err := k.Load(kconfmap.Provider(defaults, delim), nil); err != nil {
			return fmt.Errorf("load defaults for %s: %w", key, err)
		}
	}

	// 2. 从主配置加载（中等优先级，会覆盖默认值）
	mainConfig := c.kConfig.Get(key)
	if mainConfig != nil {
		// 将 mainConfig 转换为 map[string]interface{}
		if mainMap, ok := mainConfig.(map[string]interface{}); ok {
			if err := k.Load(kconfmap.Provider(mainMap, delim), nil); err != nil {
				return fmt.Errorf("load main config for %s: %w", key, err)
			}
		}
	}

	// 3. 从子配置文件加载（高优先级，会覆盖主配置）
	// 仅在 folder 或 deploy 模式下生效
	if c.kEnvStruct.ConfigMode == contract.ConfigModeFolder || c.kEnvStruct.ConfigMode == contract.ConfigModeDeploy {
		subConfigFile := filepath.Join(c.folder, key+".yaml")
		if fileExists(subConfigFile) {
			if err := k.Load(kfile.Provider(subConfigFile), kyaml.Parser()); err != nil {
				return fmt.Errorf("load sub config file for %s: %w", key, err)
			}
		}
	}

	// 4. 从环境变量加载（最高优先级）
	// 读取 APP_{KEY}_{FIELD} 格式的环境变量（如 APP_DEBUG, APP_VERSION）
	envPrefix := strings.ToUpper(key) + "_"
	for _, envPair := range os.Environ() {
		if strings.HasPrefix(envPair, envPrefix) {
			parts := strings.SplitN(envPair, "=", 2)
			if len(parts) == 2 {
				fieldKey := strings.ToLower(parts[0][len(envPrefix):])
				k.Set(fieldKey, parts[1])
			}
		}
	}

	// 5. 将配置反序列化为结构体并验证
	configStruct := config.ConfigStruct()
	if configStruct != nil {
		if err := k.UnmarshalWithConf("", configStruct, koanf.UnmarshalConf{Tag: "koanf"}); err != nil {
			return fmt.Errorf("unmarshal config struct for %s: %w", key, err)
		}

		// 调用 Validate 方法验证配置
		if err := config.Validate(configStruct); err != nil {
			return fmt.Errorf("validate config for %s: %w", key, err)
		}
	}

	c.kSubConfig[key] = k
	return nil
}

// fileExists 检查文件是否存在
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// GetSubConfig 获取子配置
// 支持两种格式：
//   - 模块名："app"
//   - 契约key："gob:app"
//
// 返回对应的子配置 Koanf 实例，如果不存在返回 nil
func (c *ConfigService) GetSubConfig(key string) *koanf.Koanf {
	c.lock.RLock()
	defer c.lock.RUnlock()

	// 如果 key 包含冒号，提取模块名
	// 例如："gob:app" -> "app"
	if strings.Contains(key, ":") {
		parts := strings.SplitN(key, ":", 2)
		if len(parts) == 2 && parts[1] != "" {
			key = parts[1]
		}
	}

	return c.kSubConfig[key]
}

// RegisterSubConfigBySubConfigRegistry 通过 SubConfigRegistry 接口注册子配置
// 这是避免循环依赖的内部方法，实际调用 RegisterSubConfig
func (c *ConfigService) RegisterSubConfigBySubConfigRegistry(cfg interface{}) error {
	if serviceConfig, ok := cfg.(framework.ServiceConfig); ok {
		return c.RegisterSubConfig(serviceConfig)
	}
	return fmt.Errorf("config must implement framework.ServiceConfig interface")
}
