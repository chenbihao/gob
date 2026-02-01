package framework

import (
	"errors"
	"sync"
	"testing"
)

// TestServiceProvider 用于测试的服务提供者
type TestServiceProvider struct {
	name        string
	isDefer     bool
	bootCalled  bool
	registered  bool
	params      []interface{}
	container   Container
	shouldError bool
}

// 实现 ServiceProvider 接口
func (sp *TestServiceProvider) Name() string {
	return sp.name
}

func (sp *TestServiceProvider) IsDefer() bool {
	return sp.isDefer
}

func (sp *TestServiceProvider) Boot(c Container) error {
	sp.container = c
	sp.bootCalled = true
	if sp.shouldError {
		return errors.New("boot error")
	}
	return nil
}

func (sp *TestServiceProvider) Params(c Container) []interface{} {
	if sp.params != nil {
		return sp.params
	}
	return []interface{}{c}
}

func (sp *TestServiceProvider) Register(c Container) NewInstance {
	sp.registered = true
	return func(params ...interface{}) (interface{}, error) {
		if sp.shouldError {
			return nil, errors.New("register error")
		}
		return &TestService{
			name: sp.name,
			data: params,
		}, nil
	}
}

// TestService 用于测试的服务
type TestService struct {
	name string
	data []interface{}
}

func (s *TestService) Name() string {
	return s.name
}

// TestNewGobContainer 测试创建新容器
func TestNewGobContainer(t *testing.T) {
	container := NewGobContainer()
	if container == nil {
		t.Fatal("NewGobContainer() should not return nil")
	}

	if container.providers == nil {
		t.Error("providers map should be initialized")
	}

	if container.instances == nil {
		t.Error("instances map should be initialized")
	}
}

// TestBind 测试 Bind 功能
func TestBind(t *testing.T) {
	tests := []struct {
		name        string
		provider    *TestServiceProvider
		expectError bool
		isDeferCase bool
	}{
		{
			name: "bind non-defer provider",
			provider: &TestServiceProvider{
				name:    "test-service",
				isDefer: false,
			},
			expectError: false,
		},
		{
			name: "bind defer provider",
			provider: &TestServiceProvider{
				name:    "defer-service",
				isDefer: true,
			},
			expectError: false,
			isDeferCase: true, // 标记这是延迟实例化服务
		},
		{
			name: "bind provider with boot error",
			provider: &TestServiceProvider{
				name:        "error-service",
				isDefer:     false,
				shouldError: true,
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := NewGobContainer()
			err := container.Bind(tt.provider)

			if tt.expectError {
				if err == nil {
					t.Error("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				// 验证服务提供者已注册
				if !container.IsBind(tt.provider.Name()) {
					t.Error("service provider should be bound")
				}

				// 验证 Boot 和 Register 只在非延迟实例化服务时被调用
				if !tt.isDeferCase {
					if !tt.provider.bootCalled {
						t.Error("Boot should be called for non-defer service")
					}
					if !tt.provider.registered {
						t.Error("Register should be called for non-defer service")
					}
				} else {
					if tt.provider.bootCalled {
						t.Error("Boot should NOT be called for defer service during binding")
					}
					if tt.provider.registered {
						t.Error("Register should NOT be called for defer service during binding")
					}
				}
			}
		})
	}
}

// TestMustBind 测试 MustBind 功能
func TestMustBind(t *testing.T) {
	t.Run("successful bind", func(t *testing.T) {
		container := NewGobContainer()
		provider := &TestServiceProvider{
			name:    "test-service",
			isDefer: false,
		}

		// 不应该 panic
		container.MustBind(provider)

		if !container.IsBind(provider.Name()) {
			t.Error("service provider should be bound")
		}
	})

	t.Run("panic on error", func(t *testing.T) {
		container := NewGobContainer()
		provider := &TestServiceProvider{
			name:        "error-service",
			isDefer:     false,
			shouldError: true,
		}

		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic but did not get one")
			}
		}()

		container.MustBind(provider)
	})
}

// TestIsBind 测试 IsBind 功能
func TestIsBind(t *testing.T) {
	container := NewGobContainer()
	provider := &TestServiceProvider{
		name:    "test-service",
		isDefer: false,
	}

	// 绑定前应该是 false
	if container.IsBind("test-service") {
		t.Error("service should not be bound before binding")
	}

	container.Bind(provider)

	// 绑定后应该是 true
	if !container.IsBind("test-service") {
		t.Error("service should be bound after binding")
	}

	// 不存在的服务应该是 false
	if container.IsBind("non-existent") {
		t.Error("non-existent service should not be bound")
	}
}

// TestMake 测试 Make 功能（单例模式）
func TestMake(t *testing.T) {
	t.Run("get existing service", func(t *testing.T) {
		container := NewGobContainer()
		provider := &TestServiceProvider{
			name:    "test-service",
			isDefer: false,
		}

		container.Bind(provider)

		// 第一次获取
		service1, err := container.Make("test-service")
		if err != nil {
			t.Fatalf("failed to get service: %v", err)
		}

		// 第二次获取
		service2, err := container.Make("test-service")
		if err != nil {
			t.Fatalf("failed to get service: %v", err)
		}

		// 验证是同一个实例（单例）
		if service1 != service2 {
			t.Error("Make should return the same instance")
		}
	})

	t.Run("get defer service", func(t *testing.T) {
		container := NewGobContainer()
		provider := &TestServiceProvider{
			name:    "defer-service",
			isDefer: true,
		}

		container.Bind(provider)

		// Bind 延迟实例化服务时，Boot 和 Register 不应该被调用
		if provider.bootCalled {
			t.Error("Boot should NOT be called during Bind for defer service")
		}
		if provider.registered {
			t.Error("Register should NOT be called during Bind for defer service")
		}

		// 第一次获取
		service1, err := container.Make("defer-service")
		if err != nil {
			t.Fatalf("failed to get defer service: %v", err)
		}

		// Make 时 Boot 和 Register 应该被调用
		if !provider.bootCalled {
			t.Error("Boot should be called for defer service on Make")
		}
		if !provider.registered {
			t.Error("Register should be called for defer service on Make")
		}

		// 第二次获取
		service2, err := container.Make("defer-service")
		if err != nil {
			t.Fatalf("failed to get defer service: %v", err)
		}

		// 验证是同一个实例（单例）
		if service1 != service2 {
			t.Error("Make should return the same instance for defer service")
		}
	})

	t.Run("get non-existent service", func(t *testing.T) {
		container := NewGobContainer()
		_, err := container.Make("non-existent")

		if err == nil {
			t.Error("expected error for non-existent service")
		}
	})
}

// TestMustMake 测试 MustMake 功能
func TestMustMake(t *testing.T) {
	t.Run("successful get", func(t *testing.T) {
		container := NewGobContainer()
		provider := &TestServiceProvider{
			name:    "test-service",
			isDefer: false,
		}

		container.Bind(provider)
		service := container.MustMake("test-service")

		if service == nil {
			t.Error("MustMake should return a service")
		}
	})

	t.Run("panic on non-existent service", func(t *testing.T) {
		container := NewGobContainer()

		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic but did not get one")
			}
		}()

		container.MustMake("non-existent")
	})
}

// TestMakeNew 测试 MakeNew 功能（每次新建实例）
func TestMakeNew(t *testing.T) {
	container := NewGobContainer()
	provider := &TestServiceProvider{
		name:    "test-service",
		isDefer: false,
	}

	container.Bind(provider)

	// 第一次获取新实例
	service1, err := container.MakeNew("test-service", nil)
	if err != nil {
		t.Fatalf("failed to get new service: %v", err)
	}

	// 第二次获取新实例
	service2, err := container.MakeNew("test-service", nil)
	if err != nil {
		t.Fatalf("failed to get new service: %v", err)
	}

	// 验证不是同一个实例
	if service1 == service2 {
		t.Error("MakeNew should return different instances")
	}

	// 验证单例实例不受影响
	service3, _ := container.Make("test-service")
	if service3 == service1 || service3 == service2 {
		t.Error("MakeNew should not affect singleton instance")
	}
}

// TestMakeNewWithParams 测试 MakeNew 使用自定义参数
func TestMakeNewWithParams(t *testing.T) {
	container := NewGobContainer()
	provider := &TestServiceProvider{
		name:   "test-service",
		params: []interface{}{"param1", "param2"},
	}

	container.Bind(provider)

	// 使用自定义参数
	service1, err := container.MakeNew("test-service", []interface{}{"custom1", "custom2"})
	if err != nil {
		t.Fatalf("failed to get new service: %v", err)
	}

	testService1 := service1.(*TestService)
	if len(testService1.data) != 2 {
		t.Errorf("expected 2 params, got %d", len(testService1.data))
	}

	// 验证使用的是自定义参数
	if testService1.data[0] != "custom1" || testService1.data[1] != "custom2" {
		t.Error("MakeNew should use custom params")
	}
}

// TestNameList 测试 NameList 功能
func TestNameList(t *testing.T) {
	container := NewGobContainer()

	// 初始应该是空列表
	names := container.NameList()
	if len(names) != 0 {
		t.Errorf("expected empty name list, got %d", len(names))
	}

	// 绑定多个服务
	container.Bind(&TestServiceProvider{name: "service1", isDefer: false})
	container.Bind(&TestServiceProvider{name: "service2", isDefer: true})
	container.Bind(&TestServiceProvider{name: "service3", isDefer: false})

	// 验证列表
	names = container.NameList()
	if len(names) != 3 {
		t.Errorf("expected 3 names, got %d", len(names))
	}

	nameSet := make(map[string]bool)
	for _, name := range names {
		nameSet[name] = true
	}

	for _, expected := range []string{"service1", "service2", "service3"} {
		if !nameSet[expected] {
			t.Errorf("expected name %s not found", expected)
		}
	}
}

// TestConcurrentAccess 测试并发安全性
func TestConcurrentAccess(t *testing.T) {
	container := NewGobContainer()

	// 绑定一个服务
	provider := &TestServiceProvider{
		name:    "test-service",
		isDefer: false,
	}
	container.Bind(provider)

	// 并发读取
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := container.Make("test-service")
			if err != nil {
				t.Errorf("concurrent get failed: %v", err)
			}
		}()
	}

	wg.Wait()
}

// TestConcurrentBind 测试并发绑定
func TestConcurrentBind(t *testing.T) {
	container := NewGobContainer()

	// 并发绑定多个服务
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			provider := &TestServiceProvider{
				name:    "service-" + string(rune(idx)),
				isDefer: false,
			}
			err := container.Bind(provider)
			if err != nil {
				t.Errorf("concurrent bind failed: %v", err)
			}
		}(i)
	}

	wg.Wait()

	// 验证所有服务都已绑定
	names := container.NameList()
	if len(names) != 10 {
		t.Errorf("expected 10 services, got %d", len(names))
	}
}

// TestProviderParams 测试服务提供者参数传递
func TestProviderParams(t *testing.T) {
	container := NewGobContainer()
	provider := &TestServiceProvider{
		name:   "test-service",
		params: []interface{}{"param1", 123, true},
	}

	container.Bind(provider)

	service, err := container.Make("test-service")
	if err != nil {
		t.Fatalf("failed to get service: %v", err)
	}

	testService := service.(*TestService)
	if len(testService.data) != 3 {
		t.Errorf("expected 3 params, got %d", len(testService.data))
	}

	if testService.data[0] != "param1" {
		t.Errorf("expected param1, got %v", testService.data[0])
	}

	if testService.data[1] != 123 {
		t.Errorf("expected 123, got %v", testService.data[1])
	}

	if testService.data[2] != true {
		t.Errorf("expected true, got %v", testService.data[2])
	}
}

// TestBootContainerAccess 测试 Boot 时访问容器
func TestBootContainerAccess(t *testing.T) {
	container := NewGobContainer()
	provider := &TestServiceProvider{
		name:    "test-service",
		isDefer: false,
	}

	container.Bind(provider)

	// 验证 Boot 时容器被正确传入
	if provider.container != container {
		t.Error("Boot should receive the container parameter")
	}
}
