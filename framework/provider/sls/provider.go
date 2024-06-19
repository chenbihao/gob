package sls

import (
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

type SLSProvider struct {
}

func (provider *SLSProvider) Name() string {
	return contract.SLSKey
}

func (provider *SLSProvider) Register(c framework.Container) framework.NewInstance {
	return NewSLSService
}

func (provider *SLSProvider) IsDefer() bool {
	return true
}

func (provider *SLSProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (provider *SLSProvider) Boot(c framework.Container) error {
	return nil
}
