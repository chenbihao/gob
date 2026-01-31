package contract

/*
## 服务介绍：
提供基础的 app 框架目录结构获取功能
## 支持命令：
[app](../command/app)
## 可初始化参数：
BaseFolder：`container.Bind(&app.AppProvider{"/workspace/gobxxx"})`
## 支持配置：无
*/

// AppKey 定义字符串凭证
const AppKey = "gob:app"

// App 定义接口（提供了获取框架相关内容，例如获取框架约定的相关目录）
type App interface {
	// AppID 表示当前这个app的唯一id, 可以用于分布式锁等
	AppID() string
	// Version 定义当前版本
	Version() string
	// // IsToolMode 是否纯工具运行模式
	// IsToolMode() bool

	// BaseFolder 定义项目基础地址
	BaseFolder() string
}
