package orm

// ServiceProvider 实现文件 provider.go

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

// GormProvider 提供App的具体实现方法
type GormProvider struct {
}

var _ framework.ServiceProvider = (*GormProvider)(nil)

// Register 注册方法
func (provider *GormProvider) Register(container framework.Container) framework.NewInstance {
	return NewGormService
}

// Boot 启动调用
func (provider *GormProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
// ORM 需要延迟初始化
func (provider *GormProvider) IsDefer() bool {
	return true
}

// Params 获取初始化参数
func (provider *GormProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

// Name 获取字符串凭证
func (provider *GormProvider) Name() string {
	return contract.ORMKey
}
