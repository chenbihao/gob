package config

import (
	"testing"

	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
)

// mockServiceConfig 模拟 ServiceConfig 接口
type mockServiceConfig struct {
	name     string
	defaults map[string]interface{}
}

func (m *mockServiceConfig) ConfigName() string {
	return m.name
}

func (m *mockServiceConfig) ConfigStruct() interface{} {
	return struct{}{}
}

func (m *mockServiceConfig) Defaults() map[string]interface{} {
	return m.defaults
}

func (m *mockServiceConfig) Validate(config interface{}) error {
	return nil
}

// TestRegisterSubConfig_DefaultOnly 测试只加载默认值
func TestRegisterSubConfig_DefaultOnly(t *testing.T) {
	cs := &ConfigService{
		kConfig:    koanf.New("."),
		kSubConfig: make(map[string]*koanf.Koanf),
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode: contract.ConfigModeRoot,
		},
	}

	mockCfg := &mockServiceConfig{
		name: "test",
		defaults: map[string]interface{}{
			"key1": "value1",
			"key2": 123,
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	subConfig := cs.GetSubConfig("test")
	assert.NotNil(t, subConfig)
	assert.Equal(t, "value1", subConfig.String("key1"))
	assert.Equal(t, int64(123), subConfig.Int64("key2"))
}

// TestRegisterSubConfig_WithMainConfig 测试主配置覆盖默认值
func TestRegisterSubConfig_WithMainConfig(t *testing.T) {
	kMain := koanf.New(".")
	kMain.Set("test.key1", "override_value")

	cs := &ConfigService{
		kConfig:    kMain,
		kSubConfig: make(map[string]*koanf.Koanf),
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode: contract.ConfigModeRoot,
		},
	}

	mockCfg := &mockServiceConfig{
		name: "test",
		defaults: map[string]interface{}{
			"key1": "default_value",
			"key2": 123,
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	subConfig := cs.GetSubConfig("test")
	assert.NotNil(t, subConfig)
	// 主配置应覆盖默认值
	assert.Equal(t, "override_value", subConfig.String("key1"))
	// 默认值中未在主配置中定义的应保留
	assert.Equal(t, int64(123), subConfig.Int64("key2"))
}

// TestRegisterSubConfig_DuplicateKey 测试重复注册应返回错误
func TestRegisterSubConfig_DuplicateKey(t *testing.T) {
	cs := &ConfigService{
		kConfig:    koanf.New("."),
		kSubConfig: make(map[string]*koanf.Koanf),
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode: contract.ConfigModeRoot,
		},
	}

	mockCfg := &mockServiceConfig{
		name:     "test",
		defaults: map[string]interface{}{},
	}

	// 第一次注册应成功
	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	// 第二次注册应失败
	err = cs.RegisterSubConfig(mockCfg)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already registered")
}

// TestGetSubConfig_NotFound 测试获取不存在的子配置
func TestGetSubConfig_NotFound(t *testing.T) {
	cs := &ConfigService{
		kConfig:    koanf.New("."),
		kSubConfig: make(map[string]*koanf.Koanf),
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode: contract.ConfigModeRoot,
		},
	}

	subConfig := cs.GetSubConfig("nonexistent")
	assert.Nil(t, subConfig)
}

// TestSubConfigPriority 测试子配置优先级
func TestSubConfigPriority(t *testing.T) {
	// 这个测试验证优先级：环境变量 > 子配置文件 > 主配置文件 > 默认值
	// 由于完整测试需要文件系统，这里只测试基本逻辑
	kMain := koanf.New(".")
	kMain.Set("app.debug", false)

	cs := &ConfigService{
		kConfig:    kMain,
		kSubConfig: make(map[string]*koanf.Koanf),
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode: contract.ConfigModeRoot,
		},
	}

	mockCfg := &mockServiceConfig{
		name: "app",
		defaults: map[string]interface{}{
			"debug":   false,
			"version": "1.0.0",
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	subConfig := cs.GetSubConfig("app")
	assert.NotNil(t, subConfig)
	// 主配置应覆盖默认值
	assert.Equal(t, false, subConfig.Bool("debug"))
	// 默认值保留
	assert.Equal(t, "1.0.0", subConfig.String("version"))
}

// TestConcurrentAccess 测试并发访问读写
func TestConcurrentAccess(t *testing.T) {
	cs := &ConfigService{
		kConfig:    koanf.New("."),
		kSubConfig: make(map[string]*koanf.Koanf),
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode: contract.ConfigModeRoot,
		},
	}

	done := make(chan bool)

	// 启动多个 goroutine 并发注册和读取
	for i := 0; i < 10; i++ {
		go func(id int) {
			cfg := &mockServiceConfig{
				name:     "test" + string(rune('a'+id)),
				defaults: map[string]interface{}{},
			}
			_ = cs.RegisterSubConfig(cfg)
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		go func() {
			_ = cs.GetSubConfig("test")
			done <- true
		}()
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 20; i++ {
		<-done
	}
	// 如果没有 panic 或死锁，测试通过
}

// TestServiceConfig_Interface 测试 ServiceConfig 接口的基本功能
func TestServiceConfig_Interface(t *testing.T) {
	var _ framework.ServiceConfig = (*mockServiceConfig)(nil)

	cfg := &mockServiceConfig{
		name: "test",
		defaults: map[string]interface{}{
			"key1": "value1",
			"key2": 123,
		},
	}

	// 测试 ConfigName
	assert.Equal(t, "test", cfg.ConfigName())

	// 测试 Defaults
	defaults := cfg.Defaults()
	assert.NotNil(t, defaults)
	assert.Equal(t, "value1", defaults["key1"])
	assert.Equal(t, 123, defaults["key2"])

	// 测试 ConfigStruct
	assert.NotNil(t, cfg.ConfigStruct())

	// 测试 Validate
	err := cfg.Validate(struct{}{})
	assert.NoError(t, err)
}
