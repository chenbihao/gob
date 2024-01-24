package ssh

// ServiceProvider 实现文件 provider.go

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

type SSHProvider struct {
	c framework.Container
}

var _ framework.ServiceProvider = (*SSHProvider)(nil)

// Register 注册方法
func (provider *SSHProvider) Register(container framework.Container) framework.NewInstance {
	return NewSSHService
}

// Boot 启动调用
func (provider *SSHProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (provider *SSHProvider) IsDefer() bool {
	return true
}

// Params 获取初始化参数
func (provider *SSHProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

// Name 获取字符串凭证
func (provider *SSHProvider) Name() string {
	return contract.SSHKey
}
