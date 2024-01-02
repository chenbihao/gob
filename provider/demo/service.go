package demo

// 实现具体的服务实例 service.go

import (
	"fmt"
	"github.com/chenbihao/gob/framework"
)

// 具体的接口实例
type DemoService struct {
	// 参数
	c framework.Container
}

// var _ Service = &DemoService{} // 确保已经实现 Service 接口
var _ Service = new(DemoService) // 确保已经实现 Service 接口

// 实现接口
func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "i am foo",
	}
}

// 初始化实例的方法
func NewDemoService(params ...interface{}) (interface{}, error) {
	// 这里需要将参数展开
	c := params[0].(framework.Container)

	fmt.Println("new demo service")
	// 返回实例
	return &DemoService{c: c}, nil
}
