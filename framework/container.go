package framework

import (
	"errors"
	"fmt"
	"sync"
)

// Container 是一个服务容器，提供绑定服务和获取服务的功能
type Container interface {

	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回 error
	Bind(provider ServiceProvider) error

	// MustBind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，如果绑定失败，那么会 panic。
	MustBind(provider ServiceProvider)

	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务
	Make(key string) (interface{}, error)

	// MustMake 根据关键字凭证获取一个服务，如果未绑定服务提供者，那么会 panic。
	// 在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) interface{}

	// MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务提供者注册的启动函数和传递的 params 参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// GobContainer 是服务容器的具体实现
type GobContainer struct {
	Container                            // 强制要求 GobContainer 实现 Container 接口
	providers map[string]ServiceProvider // providers 存储注册的服务提供者，key 为字符串凭证
	instances map[string]interface{}     // instance 存储具体的实例，key 为字符串凭证
	lock      sync.RWMutex               // lock 用于锁住对容器的变更操作
}

// NewGobContainer 创建一个服务容器
func NewGobContainer() *GobContainer {
	return &GobContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// Bind 将服务容器和关键字做了绑定
func (container *GobContainer) Bind(provider ServiceProvider) error {
	// 写锁
	container.lock.Lock()
	// key 为关键字，value 为注册的 ServiceProvider
	key := provider.Name()
	container.providers[key] = provider
	container.lock.Unlock()

	// if provider is not defer
	if !provider.IsDefer() {
		if err := provider.Boot(container); err != nil {
			return err
		}
		// 实例化方法
		params := provider.Params(container)
		instance, err := provider.Register(container)(params...)
		if err != nil {
			fmt.Println("bind service provider ", key, " error: ", err)
			return errors.New(err.Error())
		}
		container.instances[key] = instance
	}
	return nil
}

// MustBind 将服务容器和关键字做了绑定
func (container *GobContainer) MustBind(provider ServiceProvider) {
	err := container.Bind(provider)
	if err != nil {
		panic(err)
	}
}

func (container *GobContainer) IsBind(key string) bool {
	return container.findServiceProvider(key) != nil
}

func (container *GobContainer) findServiceProvider(key string) ServiceProvider {
	container.lock.RLock()
	defer container.lock.RUnlock()
	if sp, ok := container.providers[key]; ok {
		return sp
	}
	return nil
}

func (container *GobContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	// force new a
	if err := sp.Boot(container); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(container)
	}
	method := sp.Register(container)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}

// Make 方式调用内部的 make 实现
func (container *GobContainer) Make(key string) (interface{}, error) {
	return container.make(key, nil, false)
}

// MustMake 方式调用内部的 make 实现
func (container *GobContainer) MustMake(key string) interface{} {
	serv, err := container.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

// MakeNew 方式使用内部的 make 初始化
func (container *GobContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return container.make(key, params, true)
}

// 真正的实例化一个服务
func (container *GobContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	container.lock.RLock()
	defer container.lock.RUnlock()
	// 查询是否已经注册了这个服务提供者，如果没有注册，则返回错误
	sp := container.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return container.newInstance(sp, params)
	}

	// 不需要强制重新实例化，如果容器中已经实例化了，那么就直接使用容器中的实例
	if ins, ok := container.instances[key]; ok {
		return ins, nil
	}

	// 容器中还未实例化，则进行一次实例化
	inst, err := container.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	container.instances[key] = inst
	return inst, nil
}

// NameList 列出容器中所有服务提供者的字符串凭证
func (container *GobContainer) NameList() []string {
	var ret []string
	for _, provider := range container.providers {
		name := provider.Name()
		ret = append(ret, name)
	}
	return ret
}
