package id

// 实现具体的服务实例 service.go

import (
	"github.com/chenbihao/gob/framework/contract"

	"github.com/rs/xid" // 全局唯一标识符（GUID）
)

// IDService 是 IDService 的具体实现
type IDService struct{}

var _ contract.ID = (*IDService)(nil)

func NewIDService(params ...interface{}) (interface{}, error) {
	return &IDService{}, nil
}

func (s *IDService) NewID() string {
	return xid.New().String()
}
