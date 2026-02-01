package config

import (
	"fmt"
	"testing"

	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/assert"
)

// ============================================================================
// 测试辅助
// ============================================================================

// mockConfigStruct 模拟配置结构体
type mockConfigStruct struct {
	RequiredField string `koanf:"required_field"`
	OptionalField string `koanf:"optional_field"`
}

// mockServiceConfig 模拟 ServiceConfig 接口
type mockServiceConfig struct {
	name     string
	defaults map[string]interface{}
}

func (m *mockServiceConfig) ConfigName() string {
	return m.name
}

func (m *mockServiceConfig) ConfigStruct() interface{} {
	return &mockConfigStruct{}
}

func (m *mockServiceConfig) Defaults() map[string]interface{} {
	return m.defaults
}

func (m *mockServiceConfig) Validate(config interface{}) error {
	cfg, ok := config.(*mockConfigStruct)
	if !ok {
		return nil
	}
	if cfg.RequiredField == "" {
		return fmt.Errorf("required_field is required")
	}
	return nil
}

// newTestConfigService 创建测试用的ConfigService
func newTestConfigService() *ConfigService {
	return &ConfigService{
		kConfig:    koanf.New("."),
		kSubConfig: make(map[string]*koanf.Koanf),
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode: contract.ConfigModeRoot,
		},
	}
}

// ============================================================================
// GetSubConfig 方法测试组
// ============================================================================

// TestRegisterSubConfig_DefaultOnly 测试只加载默认值
func TestRegisterSubConfig_DefaultOnly(t *testing.T) {
	cs := newTestConfigService()

	mockCfg := &mockServiceConfig{
		name: "test",
		defaults: map[string]interface{}{
			"required_field": "value1",
			"optional_field": 123,
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	subConfig := cs.GetSubConfig("test")
	assert.NotNil(t, subConfig)
	assert.Equal(t, "value1", subConfig.String("required_field"))
	assert.Equal(t, int64(123), subConfig.Int64("optional_field"))
}

// TestRegisterSubConfig_WithMainConfig 测试主配置覆盖默认值
func TestRegisterSubConfig_WithMainConfig(t *testing.T) {
	kMain := koanf.New(".")
	kMain.Set("test.required_field", "override_value")

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
			"required_field": "default_value",
			"optional_field": 123,
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	subConfig := cs.GetSubConfig("test")
	assert.NotNil(t, subConfig)
	// 主配置应覆盖默认值
	assert.Equal(t, "override_value", subConfig.String("required_field"))
	// 默认值中未在主配置中定义的应保留
	assert.Equal(t, int64(123), subConfig.Int64("optional_field"))
}

// TestRegisterSubConfig_DuplicateKey 测试重复注册应返回错误
func TestRegisterSubConfig_DuplicateKey(t *testing.T) {
	cs := newTestConfigService()

	mockCfg := &mockServiceConfig{
		name: "test",
		defaults: map[string]interface{}{
			"required_field": "value",
		},
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
	cs := newTestConfigService()

	subConfig := cs.GetSubConfig("nonexistent")
	assert.Nil(t, subConfig)
}

// TestGetSubConfig_ContractKey 测试使用契约key格式获取配置
func TestGetSubConfig_ContractKey(t *testing.T) {
	cs := newTestConfigService()

	mockCfg := &mockServiceConfig{
		name: "app",
		defaults: map[string]interface{}{
			"required_field": "value",
			"debug":          true,
			"version":        "1.0.0",
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	// 使用契约key格式获取
	subConfig := cs.GetSubConfig("gob:app")
	assert.NotNil(t, subConfig)
	assert.Equal(t, true, subConfig.Bool("debug"))
	assert.Equal(t, "1.0.0", subConfig.String("version"))

	// 验证向后兼容 - 模块名格式仍然有效
	subConfig2 := cs.GetSubConfig("app")
	assert.NotNil(t, subConfig2)
	assert.Equal(t, subConfig.String("version"), subConfig2.String("version"))
}

// TestGetSubConfig_EdgeCases 测试边界情况
func TestGetSubConfig_EdgeCases(t *testing.T) {
	cs := newTestConfigService()

	mockCfg := &mockServiceConfig{
		name: "test",
		defaults: map[string]interface{}{
			"required_field": "value",
			"key":            "value",
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	testCases := []struct {
		name       string
		input      string
		shouldFind bool
	}{
		{"单冒号前缀为空", ":test", true},
		{"双冒号", "gob:test:extra", false}, // 提取为 "test:extra" 不存在
		{"前缀冒号为空", "gob:", false},
		{"仅冒号", ":", false},
		{"空字符串", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			subConfig := cs.GetSubConfig(tc.input)
			if tc.shouldFind {
				assert.NotNil(t, subConfig)
				assert.Equal(t, "value", subConfig.String("key"))
			} else {
				assert.Nil(t, subConfig)
			}
		})
	}
}

// ============================================================================
// RegisterSubConfig 验证测试组
// ============================================================================

// TestRegisterSubConfig_ValidateError 测试验证失败应返回错误
func TestRegisterSubConfig_ValidateError(t *testing.T) {
	cs := newTestConfigService()

	// 缺少 required_field，验证应失败
	mockCfg := &mockServiceConfig{
		name: "app",
		defaults: map[string]interface{}{
			"optional_field": "optional",
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validate config for app")
	assert.Contains(t, err.Error(), "required_field is required")

	// 配置不应该被注册
	subConfig := cs.GetSubConfig("app")
	assert.Nil(t, subConfig)
}

// TestRegisterSubConfig_ValidateSuccess 测试验证成功
func TestRegisterSubConfig_ValidateSuccess(t *testing.T) {
	cs := newTestConfigService()

	// 包含 required_field，验证应成功
	mockCfg := &mockServiceConfig{
		name: "app",
		defaults: map[string]interface{}{
			"required_field": "value",
			"optional_field": "optional",
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	// 配置应该被注册
	subConfig := cs.GetSubConfig("app")
	assert.NotNil(t, subConfig)
	assert.Equal(t, "value", subConfig.String("required_field"))
}

// ============================================================================
// 配置优先级测试组
// ============================================================================

// TestSubConfigPriority 测试子配置优先级
func TestSubConfigPriority(t *testing.T) {
	// 这个测试验证优先级：环境变量 > 子配置文件 > 主配置文件 > 默认值
	// 由于完整测试需要文件系统，这里只测试基本逻辑
	kMain := koanf.New(".")
	kMain.Set("app.required_field", "override_value")

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
			"required_field": "default_value",
			"optional_field": "1.0.0",
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	subConfig := cs.GetSubConfig("app")
	assert.NotNil(t, subConfig)
	// 主配置应覆盖默认值
	assert.Equal(t, "override_value", subConfig.String("required_field"))
	// 默认值保留
	assert.Equal(t, "1.0.0", subConfig.String("optional_field"))
}

// ============================================================================
// 并发安全测试组
// ============================================================================

// TestConcurrentAccess 测试并发访问读写
func TestConcurrentAccess(t *testing.T) {
	cs := newTestConfigService()

	done := make(chan bool)

	// 启动多个 goroutine 并发注册和读取
	for i := 0; i < 10; i++ {
		go func(id int) {
			cfg := &mockServiceConfig{
				name: "test" + string(rune('a'+id)),
				defaults: map[string]interface{}{
					"required_field": fmt.Sprintf("value_%d", id),
				},
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

// ============================================================================
// 接口契约测试组
// ============================================================================

// TestServiceConfig_Interface 测试 ServiceConfig 接口的基本功能
func TestServiceConfig_Interface(t *testing.T) {
	var _ framework.ServiceConfig = (*mockServiceConfig)(nil)

	cfg := &mockServiceConfig{
		name: "test",
		defaults: map[string]interface{}{
			"required_field": "value1",
			"optional_field": 123,
		},
	}

	// 测试 ConfigName
	assert.Equal(t, "test", cfg.ConfigName())

	// 测试 Defaults
	defaults := cfg.Defaults()
	assert.NotNil(t, defaults)
	assert.Equal(t, "value1", defaults["required_field"])
	assert.Equal(t, 123, defaults["optional_field"])

	// 测试 ConfigStruct
	assert.NotNil(t, cfg.ConfigStruct())

	// 测试 Validate - 成功案例
	validCfg := &mockConfigStruct{RequiredField: "value"}
	err := cfg.Validate(validCfg)
	assert.NoError(t, err)

	// 测试 Validate - 失败案例
	invalidCfg := &mockConfigStruct{RequiredField: ""}
	err = cfg.Validate(invalidCfg)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required_field is required")
}
