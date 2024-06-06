package contract

import "net/http"

/*
## 服务介绍：
提供框架最核心的结构，包括 http 和 grpc 的 Engine 结构。
## 支持命令：无
## 支持配置：无
*/

// KernelKey 提供 kenel 服务凭证
const KernelKey = "gob:kernel"

// Kernel 接口提供框架最核心的结构
type Kernel interface {
	// HttpEngine http.Handler结构，作为net/http框架使用, 实际上是gin.Engine
	HttpEngine() http.Handler
}
