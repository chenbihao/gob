package contract

import (
	"fmt"
	"github.com/chenbihao/gob/framework"
	"golang.org/x/crypto/ssh"
)

/**
	服务介绍：
		ssh 服务
	支持命令：无
	支持配置：
		```ssh.yaml
			timeout: 1s
			network: tcp
			host: 192.168.159.128 	# ip地址
			port: 22 				# 端口
			username: demo 			# 用户名
			web-pwd:
			  password: "123456" 	# 密码
			web-key:
				rsa_key: "C:/Users/99452/.ssh/id_rsa_manjarovm_demo_key"
				known_hosts: "C:/Users/99452/.ssh/known_hosts"
		```
**/

const SSHKey = "gob:ssh"

// SSHService 表示一个ssh服务
type SSHService interface {
	// GetClient 获取ssh连接实例
	GetClient(option ...SSHOption) (*ssh.Client, error)
}

// SSHOption 代表初始化的时候的选项
type SSHOption func(container framework.Container, config *SSHConfig) error

// SSHConfig 为gob定义的SSH配置结构
type SSHConfig struct {
	NetWork string
	Host    string
	Port    string
	*ssh.ClientConfig
}

// UniqKey 用来唯一标识一个SSHConfig配置
func (config *SSHConfig) UniqKey() string {
	return fmt.Sprintf("%v_%v_%v", config.Host, config.Port, config.User)
}
