package framework

// ServiceConfig 定义一个服务提供者提供的服务配置项
type ServiceConfig interface {

	//// Koanf *Koanf ？
	//Koanf() koanf.Koanf

	//// ConfigStruct 配置值
	//ConfigStruct() any

	// ConfigName 模块名称，用于标识配置的唯一键
	ConfigName() string

	// ConfigStruct 返回模块配置的结构体实例
	// 该结构体会被配置服务用来解析和存储配置数据
	ConfigStruct() interface{}

	// Defaults 返回模块配置的默认值
	// 返回的 map 应该与 ConfigStruct 的结构相对应
	Defaults() map[string]interface{}

	// Validate 验证配置是否有效
	// 在配置加载后调用，确保配置符合预期
	Validate(config interface{}) error
}
