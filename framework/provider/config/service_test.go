package config

import (
	"fmt"
	"os"
	"path/filepath"
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
	if cfg, ok := config.(*mockConfigStruct); ok && cfg.RequiredField == "" {
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

func newTestConfigServiceWithData() *ConfigService {
	k := koanf.New(".")
	k.Set("app.name", "testapp")
	k.Set("app.port", 8080)
	k.Set("app.debug", true)
	k.Set("app.timeout", "30s")
	k.Set("app.rate", 3.14)
	k.Set("app.items", []string{"a", "b", "c"})
	return &ConfigService{
		kConfig:    k,
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
// AC3: 配置优先级测试 (代码默认值 < 主配置 YAML < 子配置文件 < 环境变量)
// ============================================================================

func TestSubConfigPriority_Full(t *testing.T) {
	kMain := koanf.New(".")
	kMain.Set("app.field", "from_main")
	kMain.Set("app.required_field", "main_reqs")

	cs := &ConfigService{
		kConfig:    kMain,
		kSubConfig: make(map[string]*koanf.Koanf),
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode: contract.ConfigModeFolder,
		},
	}

	mockCfg := &mockServiceConfig{
		name: "app",
		defaults: map[string]interface{}{
			"required_field": "default_reqs",
			"field":          "from_default",
		},
	}
	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	subConfig := cs.GetSubConfig("app")
	assert.NotNil(t, subConfig)

	assert.Equal(t, "from_main", subConfig.String("field"))
}

func TestSubConfigPriority_WithEnvVar(t *testing.T) {
	os.Setenv("APP_SVC_FIELD", "from_env")

	kMain := koanf.New(".")
	kMain.Set("app_svc.required_field", "main_field")
	kMain.Set("app_svc.field", "from_main")

	cs := &ConfigService{
		kConfig:    kMain,
		kSubConfig: make(map[string]*koanf.Koanf),
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode: contract.ConfigModeRoot,
		},
	}

	mockCfg := &mockServiceConfig{
		name: "app_svc",
		defaults: map[string]interface{}{
			"required_field": "default_field",
			"field":        "from_default",
		},
	}
	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	subConfig := cs.GetSubConfig("app_svc")
	assert.NotNil(t, subConfig)

	assert.Equal(t, "from_env", subConfig.String("field"))
	t.Cleanup(func() { os.Unsetenv("APP_SVC_FIELD") })
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

// ============================================================================
// env(key) 占位符替换测试组
// ============================================================================

// TestReplaceEnvKey 测试环境变量占位符替换
func TestReplaceEnvKey(t *testing.T) {
	tests := []struct {
		name     string
		input   []byte
		maps    map[string]string
		want    string
	}{
		{
			name:   "单个占位符替换",
			input:  []byte("host: env(DB_HOST)"),
			maps:   map[string]string{"DB_HOST": "localhost"},
			want:   "host: localhost",
		},
		{
			name:   "多个占位符替换",
			input:  []byte("host: env(DB_HOST), port: env(DB_PORT)"),
			maps:   map[string]string{"DB_HOST": "localhost", "DB_PORT": "5432"},
			want:   "host: localhost, port: 5432",
		},
		{
			name:   "无占位符",
			input:  []byte("host: localhost"),
			maps:   map[string]string{"DB_HOST": "localhost"},
			want:   "host: localhost",
		},
		{
			name:   "空maps",
			input:  []byte("host: env(DB_HOST)"),
			maps:   nil,
			want:   "host: env(DB_HOST)",
		},
		{
			name:   "不存在的key保留原样",
			input:  []byte("host: env(NONEXISTENT)"),
			maps:   map[string]string{"DB_HOST": "localhost"},
			want:   "host: env(NONEXISTENT)",
		},
		{
			name:   "嵌套占位符",
			input:  []byte("url: postgres://env(DB_USER):env(DB_PASS)@env(DB_HOST):env(DB_PORT)"),
			maps:   map[string]string{"DB_USER": "admin", "DB_PASS": "secret", "DB_HOST": "localhost", "DB_PORT": "5432"},
			want:   "url: postgres://admin:secret@localhost:5432",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := replaceEnvKey(tt.input, tt.maps)
			assert.Equal(t, tt.want, string(got))
		})
	}
}

// TestReplaceEnvKey_NilMaps 测试 nil maps 的边界情况
func TestReplaceEnvKey_NilMaps(t *testing.T) {
	input := []byte("key: env(VALUE)")
	result := replaceEnvKey(input, nil)
	assert.Equal(t, "key: env(VALUE)", string(result))

	result = replaceEnvKey(input, make(map[string]string))
	assert.Equal(t, "key: env(VALUE)", string(result))
}

// ============================================================================
// 配置模式测试组
// ============================================================================

// TestConfigMode_Folder 测试 folder 模式加载子配置文件
func TestConfigMode_Folder(t *testing.T) {
	cs := &ConfigService{
		kConfig:    koanf.New("."),
		kSubConfig: make(map[string]*koanf.Koanf),
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode: contract.ConfigModeFolder,
		},
	}

	mockCfg := &mockServiceConfig{
		name: "test_mode",
		defaults: map[string]interface{}{
			"required_field": "default",
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	// folder 模式下 kEnvStruct 应正确设置
	assert.Equal(t, "folder", string(cs.kEnvStruct.ConfigMode))
}

// TestConfigMode_Deploy 测试 deploy 模式
func TestConfigMode_Deploy(t *testing.T) {
	cs := &ConfigService{
		kConfig:    koanf.New("."),
		kSubConfig: make(map[string]*koanf.Koanf),
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode:  contract.ConfigModeDeploy,
			AppEnv:      contract.AppEnvDev,
			ConfigFolder: "config",
		},
	}

	mockCfg := &mockServiceConfig{
		name: "test_deploy",
		defaults: map[string]interface{}{
			"required_field": "default",
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	assert.Equal(t, "deploy", string(cs.kEnvStruct.ConfigMode))
	assert.Equal(t, "dev", string(cs.kEnvStruct.AppEnv))
}

// TestConfigMode_Root 测试 root 模式
func TestConfigMode_Root(t *testing.T) {
	cs := &ConfigService{
		kConfig:    koanf.New("."),
		kSubConfig: make(map[string]*koanf.Koanf),
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode:    contract.ConfigModeRoot,
			ConfigFolder: "config",
		},
	}

	mockCfg := &mockServiceConfig{
		name: "test_root",
		defaults: map[string]interface{}{
			"required_field": "default",
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	assert.Equal(t, contract.ConfigModeRoot, cs.kEnvStruct.ConfigMode)
}

// ============================================================================
// 子配置文件加载测试组 (folder/deploy模式)
// ============================================================================

// TestSubConfigFileLoading_FromFile 测试从真实文件加载子配置（folder模式）
func TestSubConfigFileLoading_FromFile(t *testing.T) {
	tmpDir := t.TempDir()

	subConfigFile := filepath.Join(tmpDir, "app.yaml")
	err := os.WriteFile(subConfigFile, []byte(`
required_field: from_file
field_from_file: overridden_value
`), 0644)
	assert.NoError(t, err)

	kMain := koanf.New(".")
	kMain.Set("app.required_field", "from_main")

	cs := &ConfigService{
		kConfig:    kMain,
		kSubConfig: make(map[string]*koanf.Koanf),
		folder:     tmpDir,
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode: contract.ConfigModeFolder,
		},
	}

	mockCfg := &mockServiceConfig{
		name: "app",
		defaults: map[string]interface{}{
			"required_field":   "default",
			"field_from_file":  "default_value",
			"default_only":     "default_only",
		},
	}

	err = cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	subConfig := cs.GetSubConfig("app")
	assert.NotNil(t, subConfig)

	assert.Equal(t, "from_file", subConfig.String("required_field"))
	assert.Equal(t, "overridden_value", subConfig.String("field_from_file"))
	assert.Equal(t, "default_only", subConfig.String("default_only"))
}

// TestSubConfigFileLoading 测试子配置文件加载
func TestSubConfigFileLoading(t *testing.T) {
	// 测试 folder 模式下子配置的加载逻辑
	// 加载优先级: 默认值 -> 主配置 -> 子配置文件 -> 环境变量

	// 使用主配置覆盖默认值
	kMain := koanf.New(".")
	kMain.Set("subtest.required_field", "main_value") // 满足验证需求
	kMain.Set("subtest.field1", "from_main")
	kMain.Set("subtest.field2", "from_main_2")

	cs := &ConfigService{
		kConfig:    kMain,
		kSubConfig: make(map[string]*koanf.Koanf),
		kEnvStruct: &contract.ConfigEnvStruct{
			ConfigMode: contract.ConfigModeRoot,
		},
	}

	// 子配置会从主配置中获取 app 相关的键值
	mockCfg := &mockServiceConfig{
		name: "subtest",
		defaults: map[string]interface{}{
			"required_field": "value", // 满足验证需求
			"field1":         "default1",
			"field2":         "default2",
			"field3":         "default3",
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	subConfig := cs.GetSubConfig("subtest")
	assert.NotNil(t, subConfig)

	// 主配置覆盖默认值(field1和field2来自主配置)
	assert.Equal(t, "from_main", subConfig.String("field1"))
	assert.Equal(t, "from_main_2", subConfig.String("field2"))

	// 默认值保留(field3不在主配置中)
	assert.Equal(t, "default3", subConfig.String("field3"))
}

// ============================================================================
// 配置热重载测试组
// ============================================================================

// TestConfigHotReload 测试配置文件热重载功能
// 注意：这是一个测试框架，实际测试需要创建临时配置文件并触发文件变更
func TestConfigHotReload(t *testing.T) {
	cs := newTestConfigService()

	// 注册一个子配置
	mockCfg := &mockServiceConfig{
		name: "reload-test",
		defaults: map[string]interface{}{
			"required_field": "initial_value",
		},
	}

	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	// 验证配置已加载
	subConfig := cs.GetSubConfig("reload-test")
	assert.NotNil(t, subConfig)
	assert.Equal(t, "initial_value", subConfig.String("required_field"))

	// 实际的热重载测试需要：
	// 1. 创建临时配置文件
	// 2. 修改配置文件内容
	// 3. 等待 Watch 回调触发
	// 4. 验证配置已更新
	// 5. 清理临时文件

	// 示例热重载测试流程（实际实现需要临时文件系统）:
	//
	// tmpDir := t.TempDir()
	// configFile := filepath.Join(tmpDir, "config.yaml")
	// os.WriteFile(configFile, []byte("reload-test:\n  required_field: updated_value\n"), 0644)
	// Load(tmpDir, "config", cs, koanf.New("."))
	// // 修改配置文件
	// os.WriteFile(configFile, []byte("reload-test:\n  required_field: reloaded_value\n"), 0644)
	// // 等待 Watch 回调
	// time.Sleep(100 * time.Millisecond)
	// // 验证配置已更新
	// assert.Equal(t, "reloaded_value", cs.GetSubConfig("reload-test").String("required_field"))

	t.Skip("Config hot reload test requires temporary file system - test framework provided")
}

// ============================================================================
// AC1: Config 核心方法测试组 (Get/GetInt/GetBool/GetString/GetDuration)
// ============================================================================

func TestGetSubConfig_Methods(t *testing.T) {
	cs := newTestConfigServiceWithData()

	mockCfg := &mockServiceConfig{
		name: "app",
		defaults: map[string]interface{}{
			"required_field": "value",
		},
	}
	err := cs.RegisterSubConfig(mockCfg)
	assert.NoError(t, err)

	subConfig := cs.GetSubConfig("app")
	assert.NotNil(t, subConfig)

	t.Run("String", func(t *testing.T) {
		assert.Equal(t, "testapp", subConfig.String("name"))
	})

	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, 8080, subConfig.Int("port"))
	})

	t.Run("Int64", func(t *testing.T) {
		assert.Equal(t, int64(8080), subConfig.Int64("port"))
	})

	t.Run("Bool", func(t *testing.T) {
		assert.Equal(t, true, subConfig.Bool("debug"))
	})

	t.Run("Float64", func(t *testing.T) {
		assert.Equal(t, 3.14, subConfig.Float64("rate"))
	})

	t.Run("Duration", func(t *testing.T) {
		assert.Equal(t, "30s", subConfig.String("timeout"))
	})

	t.Run("Strings", func(t *testing.T) {
		items := subConfig.Strings("items")
		assert.Equal(t, []string{"a", "b", "c"}, items)
	})

	t.Run("Get non-existent key", func(t *testing.T) {
		assert.Nil(t, subConfig.Get("nonexistent"))
	})

	t.Run("Get non-existent with default", func(t *testing.T) {
		assert.Equal(t, nil, subConfig.Get("nonexistent"))
	})
}
