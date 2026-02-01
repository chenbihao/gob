package app

import (
	"testing"

	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/provider/config"
)

// NewTestContainerWithBaseProviders 创建已注册基础服务的测试容器
// 按正确顺序注册：App（Config 需要）→ Config（其他服务需要）
// 避免每个测试重复注册逻辑和循环依赖问题
func NewTestContainerWithBaseProviders() *framework.TestContainer {
	testContainer := framework.NewTestContainer()

	// 1. 先注册 AppProvider（ConfigService 初始化需要 App.BaseFolder()）
	testContainer.MustBind(&AppProvider{})

	// 2. 再注册 ConfigProvider（其他服务可能需要 Config）
	testContainer.MustBind(&config.ConfigProvider{})

	return testContainer
}

// TestAppProvider_Name 测试 Name 方法
func TestAppProvider_Name(t *testing.T) {
	provider := &AppProvider{}
	name := provider.Name()
	if name != contract.AppKey {
		t.Errorf("expected name %s, got %s", contract.AppKey, name)
	}
}

// TestAppProvider_IsDefer 测试 IsDefer 方法
func TestAppProvider_IsDefer(t *testing.T) {
	provider := &AppProvider{}
	isDefer := provider.IsDefer()
	if isDefer != false {
		t.Errorf("expected IsDefer to be false, got %v", isDefer)
	}
}

// TestAppProvider_Boot 测试 Boot 方法
func TestAppProvider_Boot(t *testing.T) {
	// 使用辅助函数获取已注册基础服务的容器
	container := NewTestContainerWithBaseProviders().GetContainer()

	provider := &AppProvider{}
	err := provider.Boot(container)
	if err != nil {
		t.Errorf("Boot should not return error, got %v", err)
	}
}

// TestAppProvider_Params 测试 Params 方法
func TestAppProvider_Params(t *testing.T) {
	tests := []struct {
		name       string
		baseFolder string
	}{
		{
			name:       "with base folder",
			baseFolder: "/test/path",
		},
		{
			name:       "with empty base folder",
			baseFolder: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider := &AppProvider{BaseFolder: tt.baseFolder}
			container := framework.NewTestContainer().GetContainer()

			params := provider.Params(container)

			if len(params) != 2 {
				t.Errorf("expected 2 params, got %d", len(params))
			}

			// 第一个参数应该是容器
			if params[0] != container {
				t.Error("first param should be container")
			}

			// 第二个参数应该是 baseFolder
			if params[1] != tt.baseFolder {
				t.Errorf("expected baseFolder %s, got %v", tt.baseFolder, params[1])
			}
		})
	}
}

// TestAppProvider_Register 测试 Register 方法
func TestAppProvider_Register(t *testing.T) {
	provider := &AppProvider{}
	container := framework.NewTestContainer().GetContainer()

	newInstance := provider.Register(container)
	if newInstance == nil {
		t.Error("Register should return a NewInstance function")
	}

	// 验证返回的函数可以正确创建实例
	instance, err := newInstance(container, "")
	if err != nil {
		t.Errorf("NewInstance should not return error, got %v", err)
	}

	if instance == nil {
		t.Error("NewInstance should return an instance")
	}

	// 验证实例类型
	appService, ok := instance.(*AppService)
	if !ok {
		t.Error("NewInstance should return *AppService type")
	}

	if appService == nil {
		t.Error("appService should not be nil")
	}
}

// TestAppProvider_RegisterWithWrongParams 测试使用错误参数调用 Register 返回的函数
func TestAppProvider_RegisterWithWrongParams(t *testing.T) {
	provider := &AppProvider{}
	container := framework.NewTestContainer().GetContainer()

	newInstance := provider.Register(container)

	// 传入错误的参数数量
	_, err := newInstance(container)
	if err == nil {
		t.Error("expected error with wrong number of params")
	}

	// 传入过多参数
	_, err = newInstance(container, "", "extra")
	if err == nil {
		t.Error("expected error with too many params")
	}
}

// TestAppProvider_FullLifecycle 测试完整的服务提供者生命周期
func TestAppProvider_FullLifecycle(t *testing.T) {
	tests := []struct {
		name       string
		baseFolder string
	}{
		{
			name:       "with custom base folder",
			baseFolder: "/workspace/gob-project",
		},
		{
			name:       "with empty base folder",
			baseFolder: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := framework.NewTestContainer()
			provider := &AppProvider{BaseFolder: tt.baseFolder}

			// 绑定服务
			err := container.Bind(provider)
			if err != nil {
				t.Fatalf("Bind failed: %v", err)
			}

			// 验证服务已绑定
			if !container.IsBind(contract.AppKey) {
				t.Error("service should be bound")
			}

			// 获取服务
			app := container.MustMake(contract.AppKey).(*AppService)

			// 验证服务属性
			if app == nil {
				t.Fatal("app service should not be nil")
			}

			if app.container != container.GetContainer() {
				t.Error("app.container should be the container instance")
			}

			if app.appID == "" {
				t.Error("appID should not be empty")
			}

			// 验证 baseFolder（如果设置了）
			if tt.baseFolder != "" && app.baseFolder != tt.baseFolder {
				t.Errorf("expected baseFolder %s, got %s", tt.baseFolder, app.baseFolder)
			}
		})
	}
}

// TestAppProvider_Contract 实现 ServiceProvider 接口
func TestAppProvider_Contract(t *testing.T) {
	var _ framework.ServiceProvider = (*AppProvider)(nil)
}

// BenchmarkAppProvider_Boot 性能测试
func BenchmarkAppProvider_Boot(b *testing.B) {
	provider := &AppProvider{}
	// 使用辅助函数获取已注册基础服务的容器
	container := NewTestContainerWithBaseProviders().GetContainer()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = provider.Boot(container)
	}
}
