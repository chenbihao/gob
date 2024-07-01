package env

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

type GobTestingEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *GobTestingEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewTestingEnv
}

// Boot will called when the service instantiate
func (provider *GobTestingEnvProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *GobTestingEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *GobTestingEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

// / Name define the name for this service
func (provider *GobTestingEnvProvider) Name() string {
	return contract.EnvKey
}
