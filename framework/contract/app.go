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

	// LoadAppConfig 加载新的AppConfig，key为对应的函数转为小写下划线，比如ConfigFolder => config_folder
	LoadAppConfig(kv map[string]string)

	// ---------------- 根目录 ----------------

	// BaseFolder 定义项目基础地址
	BaseFolder() string

	// ---------------- 根目录下 ----------------

	// AppFolder 定义业务代码所在的目录，用于监控文件变更使用
	AppFolder() string

	// ConfigFolder 定义了配置文件的路径
	ConfigFolder() string
	// StorageFolder 存储文件地址
	StorageFolder() string

	// // TestFolder 存放测试所需要的信息
	// TestFolder() string

	// ---------------- app 目录下 ----------------

	// ProviderFolder 定义业务自己的通用服务提供者地址
	ProviderFolder() string

	// // ConsoleFolderr 定义业务自己的命令行服务提供者地址
	// CommandFolder() string
	// // WailsFolder 定义业务自己的 app 服务提供者地址
	// WailsFolder() string

	// // HttpFolderr 定义业务自己的 web 服务提供者地址
	// HttpFolder() string
	// // MiddlewareFolder 定义业务自己定义的中间件
	// MiddlewareFolder() string

	// ---------------- storage 目录下 ----------------

	// LogFolder 定义了日志所在路径
	LogFolder() string
	// RuntimeFolder 定义业务的运行中间态信息
	RuntimeFolder() string
}
