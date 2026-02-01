package framework

import "sync"

// MockServiceProvider 用于测试的模拟服务提供者
type MockServiceProvider struct {
	name         string
	isDefer      bool
	bootFunc     func(c Container) error
	paramsFunc   func(c Container) []interface{}
	registerFunc func(c Container) NewInstance
}

// NewMockServiceProvider 创建模拟服务提供者
func NewMockServiceProvider(name string, opts ...MockOption) *MockServiceProvider {
	sp := &MockServiceProvider{
		name:       name,
		isDefer:    false,
		bootFunc:   func(c Container) error { return nil },
		paramsFunc: func(c Container) []interface{} { return nil },
		registerFunc: func(c Container) NewInstance {
			return func(params ...interface{}) (interface{}, error) {
				return &MockService{name: name}, nil
			}
		},
	}

	for _, opt := range opts {
		opt(sp)
	}

	return sp
}

// MockOption 配置选项
type MockOption func(*MockServiceProvider)

// WithDefer 设置为延迟实例化
func WithDefer(isDefer bool) MockOption {
	return func(sp *MockServiceProvider) {
		sp.isDefer = isDefer
	}
}

// WithBootFunc 设置 Boot 函数
func WithBootFunc(fn func(c Container) error) MockOption {
	return func(sp *MockServiceProvider) {
		sp.bootFunc = fn
	}
}

// WithParamsFunc 设置 Params 函数
func WithParamsFunc(fn func(c Container) []interface{}) MockOption {
	return func(sp *MockServiceProvider) {
		sp.paramsFunc = fn
	}
}

// WithRegisterFunc 设置 Register 函数
func WithRegisterFunc(fn func(c Container) NewInstance) MockOption {
	return func(sp *MockServiceProvider) {
		sp.registerFunc = fn
	}
}

// 实现 ServiceProvider 接口
func (sp *MockServiceProvider) Name() string {
	return sp.name
}

func (sp *MockServiceProvider) IsDefer() bool {
	return sp.isDefer
}

func (sp *MockServiceProvider) Boot(c Container) error {
	return sp.bootFunc(c)
}

func (sp *MockServiceProvider) Params(c Container) []interface{} {
	return sp.paramsFunc(c)
}

func (sp *MockServiceProvider) Register(c Container) NewInstance {
	return sp.registerFunc(c)
}

// MockService 用于测试的模拟服务
type MockService struct {
	name string
	data interface{}
}

func (s *MockService) Name() string {
	return s.name
}

// TestContainer 测试容器辅助类
type TestContainer struct {
	container Container
}

// NewTestContainer 创建测试容器
func NewTestContainer() *TestContainer {
	return &TestContainer{
		container: NewGobContainer(),
	}
}

// Bind 绑定服务
func (tc *TestContainer) Bind(provider ServiceProvider) error {
	return tc.container.Bind(provider)
}

// MustBind 强制绑定服务（失败会 panic）
func (tc *TestContainer) MustBind(provider ServiceProvider) {
	tc.container.MustBind(provider)
}

// Make 获取服务
func (tc *TestContainer) Make(key string) (interface{}, error) {
	return tc.container.Make(key)
}

// MustMake 强制获取服务（失败会 panic）
func (tc *TestContainer) MustMake(key string) interface{} {
	return tc.container.MustMake(key)
}

// MakeNew 获取新实例
func (tc *TestContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return tc.container.MakeNew(key, params)
}

// IsBind 检查是否已绑定
func (tc *TestContainer) IsBind(key string) bool {
	return tc.container.IsBind(key)
}

// NameList 获取服务列表
func (tc *TestContainer) NameList() []string {
	gobContainer, ok := tc.container.(*GobContainer)
	if !ok {
		return []string{}
	}
	return gobContainer.NameList()
}

// GetContainer 获取底层容器
func (tc *TestContainer) GetContainer() Container {
	return tc.container
}

// RunConcurrentTest 并发测试辅助函数
func RunConcurrentTest(t interface{}, times int, fn func(int)) {
	var wg sync.WaitGroup
	for i := 0; i < times; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fn(idx)
		}(i)
	}
	wg.Wait()
}
