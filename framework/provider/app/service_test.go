package app

import (
	"testing"

	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

// TestNewGobApp 测试 NewGobApp 函数
func TestNewGobApp(t *testing.T) {
	tests := []struct {
		name      string
		params    []any
		expectErr bool
		checkFunc func(*testing.T, interface{})
	}{
		{
			name: "valid params with container and baseFolder",
			params: []any{
				framework.NewTestContainer().GetContainer(),
				"/test/path",
			},
			expectErr: false,
			checkFunc: func(t *testing.T, instance interface{}) {
				app, ok := instance.(*AppService)
				if !ok {
					t.Error("instance should be *AppService type")
				}
				if app == nil {
					t.Fatal("app should not be nil")
				}
				if app.appID == "" {
					t.Error("appID should not be empty")
				}
				if app.baseFolder != "/test/path" {
					t.Errorf("expected baseFolder /test/path, got %s", app.baseFolder)
				}
			},
		},
		{
			name: "valid params with empty baseFolder",
			params: []any{
				framework.NewTestContainer().GetContainer(),
				"",
			},
			expectErr: false,
			checkFunc: func(t *testing.T, instance interface{}) {
				app, ok := instance.(*AppService)
				if !ok {
					t.Error("instance should be *AppService type")
				}
				if app == nil {
					t.Fatal("app should not be nil")
				}
				if app.baseFolder != "" {
					t.Errorf("expected empty baseFolder, got %s", app.baseFolder)
				}
			},
		},
		{
			name:      "missing params",
			params:    []any{},
			expectErr: true,
			checkFunc: nil,
		},
		{
			name:      "only one param",
			params:    []any{framework.NewTestContainer().GetContainer()},
			expectErr: true,
			checkFunc: nil,
		},
		{
			name:      "too many params",
			params:    []any{framework.NewTestContainer().GetContainer(), "/path", "extra"},
			expectErr: true,
			checkFunc: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instance, err := NewGobApp(tt.params...)

			if tt.expectErr {
				if err == nil {
					t.Error("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if tt.checkFunc != nil {
					tt.checkFunc(t, instance)
				}
			}
		})
	}
}

// TestAppService_AppID 测试 AppID 方法
func TestAppService_AppID(t *testing.T) {
	container := framework.NewTestContainer()
	provider := &AppProvider{BaseFolder: "/test/path"}
	container.MustBind(provider)

	app := container.MustMake(contract.AppKey).(*AppService)

	appID := app.AppID()
	if appID == "" {
		t.Error("AppID should not be empty")
	}

	// 验证每次获取都是同一个实例（单例）
	app2 := container.MustMake(contract.AppKey).(*AppService)
	if app2.AppID() != appID {
		t.Error("AppID should be the same for singleton instance")
	}

	// 验证不同的实例有不同的 ID（通过 MakeNew）
	app3, _ := container.MakeNew(contract.AppKey, nil)
	if app3.(*AppService).AppID() == appID {
		t.Error("MakeNew should create instance with different AppID")
	}
}

// TestAppService_Version 测试 Version 方法
func TestAppService_Version(t *testing.T) {
	container := framework.NewTestContainer()
	provider := &AppProvider{}
	container.MustBind(provider)

	app := container.MustMake(contract.AppKey).(*AppService)

	version := app.Version()
	if version == "" {
		t.Error("Version should not be empty")
	}

	if version != framework.Version {
		t.Errorf("expected version %s, got %s", framework.Version, version)
	}
}

// TestAppService_BaseFolder 测试 BaseFolder 方法
func TestAppService_BaseFolder(t *testing.T) {
	tests := []struct {
		name          string
		setBaseFolder string
		expectCustom  bool
	}{
		{
			name:          "with custom base folder",
			setBaseFolder: "/workspace/test",
			expectCustom:  true,
		},
		{
			name:          "without custom base folder",
			setBaseFolder: "",
			expectCustom:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := framework.NewTestContainer()
			provider := &AppProvider{BaseFolder: tt.setBaseFolder}
			container.MustBind(provider)

			app := container.MustMake(contract.AppKey).(*AppService)

			baseFolder := app.BaseFolder()
			if baseFolder == "" {
				t.Error("BaseFolder should not be empty")
			}

			if tt.expectCustom {
				if baseFolder != tt.setBaseFolder {
					t.Errorf("expected BaseFolder %s, got %s", tt.setBaseFolder, baseFolder)
				}
			}
			// 没有设置自定义 baseFolder 时，会使用 util.GetExecDirectory()，这里不做具体值验证
		})
	}
}

// TestAppService_Container 测试容器引用
func TestAppService_Container(t *testing.T) {
	container := framework.NewTestContainer()
	testContainer := container.GetContainer()

	provider := &AppProvider{}
	container.MustBind(provider)

	app := container.MustMake(contract.AppKey).(*AppService)

	if app.container != testContainer {
		t.Error("app.container should reference the container")
	}
}

// TestAppService_Contract 实现 contract.App 接口
func TestAppService_Contract(t *testing.T) {
	var _ contract.App = (*AppService)(nil)
}

// TestAppService_MultipleInstances 测试多个实例
func TestAppService_MultipleInstances(t *testing.T) {
	container := framework.NewTestContainer()
	provider := &AppProvider{BaseFolder: "/test/path"}
	container.MustBind(provider)

	// 获取多个实例，验证它们是同一个（单例）
	app1 := container.MustMake(contract.AppKey).(*AppService)
	app2 := container.MustMake(contract.AppKey).(*AppService)
	app3 := container.MustMake(contract.AppKey).(*AppService)

	if app1 != app2 || app2 != app3 || app1 != app3 {
		t.Error("Make should return the same instance (singleton)")
	}

	// 使用 MakeNew 获取新实例
	app4, _ := container.MakeNew(contract.AppKey, nil)
	if app4.(*AppService) == app1 {
		t.Error("MakeNew should return a different instance")
	}
}

// TestAppService_ConcurrentAccess 测试并发访问
func TestAppService_ConcurrentAccess(t *testing.T) {
	container := framework.NewTestContainer()
	provider := &AppProvider{}
	container.MustBind(provider)

	app := container.MustMake(contract.AppKey).(*AppService)

	// 并发调用各种方法
	framework.RunConcurrentTest(t, 100, func(i int) {
		_ = app.AppID()
		_ = app.Version()
		_ = app.BaseFolder()
	})
}

// TestAppService_Integration 集成测试
func TestAppService_Integration(t *testing.T) {
	container := framework.NewTestContainer()
	provider := &AppProvider{BaseFolder: "/integration/test"}

	// 绑定服务
	err := container.Bind(provider)
	if err != nil {
		t.Fatalf("Bind failed: %v", err)
	}

	// 验证绑定
	if !container.IsBind(contract.AppKey) {
		t.Error("service should be bound")
	}

	// 获取服务
	app, err := container.Make(contract.AppKey)
	if err != nil {
		t.Fatalf("Make failed: %v", err)
	}

	appService := app.(*AppService)

	// 验证所有方法
	appID := appService.AppID()
	if appID == "" {
		t.Error("AppID should not be empty")
	}

	version := appService.Version()
	if version != framework.Version {
		t.Errorf("Version mismatch: expected %s, got %s", framework.Version, version)
	}

	baseFolder := appService.BaseFolder()
	if baseFolder == "" {
		t.Error("BaseFolder should not be empty")
	}

	// 验证 NameList 包含 AppKey
	names := container.NameList()
	found := false
	for _, name := range names {
		if name == contract.AppKey {
			found = true
			break
		}
	}
	if !found {
		t.Error("AppKey should be in NameList")
	}
}

// BenchmarkNewGobApp 性能测试
func BenchmarkNewGobApp(b *testing.B) {
	container := framework.NewTestContainer().GetContainer()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewGobApp(container, "/test/path")
	}
}

// BenchmarkAppService_AppID 性能测试
func BenchmarkAppService_AppID(b *testing.B) {
	container := framework.NewTestContainer()
	provider := &AppProvider{}
	container.MustBind(provider)

	app := container.MustMake(contract.AppKey).(*AppService)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = app.AppID()
	}
}
