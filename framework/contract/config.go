package contract

/*
## 服务介绍：
提供基础的配置文件获取方法
## 支持命令：
[config](../command/config)
## 支持配置：
[config](../config/config)
*/

// ConfigKey 是配置服务字符串凭证
const ConfigKey = "gob:config"

// Config 定义了配置文件服务，读取配置文件，支持点分割的路径读取
// 例如: .Get("app.name") 表示从 app 文件中读取 name 属性
// 建议使用 yaml 属性, https://yaml.org/spec/1.2/spec.html
type Config interface {
}
