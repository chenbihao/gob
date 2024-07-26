package contract

import (
	"context"
	"time"
)

/*
## 服务介绍：

提供日志记录相关操作，目前支持控制台、单文件、切割文件、阿里云SLS 等日志输出通道

> 阿里云SLS 需要配置对应的 sls 服务 [sls](./sls)

## 支持命令：无
## 支持配置：

通过配置文件 `config/[env]/log.yaml` 可以配置缓存服务的驱动和参数，如下是一个配置示例：

> 注意：rotate库已于21年存档，寻找替代品准备移除

```yaml
## drivers（启用的日志通道）默认值: [console]
## level（日志等级）默认值: info
## formatter（格式化器）默认值: text
## folder（文件目录）默认值取 app 的 LogFolder地址
## file（文件名）默认值：app.log

# 启用的日志通道
drivers: [console,single,rotate,rolling]
level: trace
formatter: text
folder: ./storage/log/

# 控制台输出
console:
  level: info
  formatter: text

# 单文件输出
single:
  level: info
  folder: ./storage/log/
  file: app_single.log

# 按日期切割（rotate库已于21年存档，寻找替代品准备移除）
rotate:
  level: info               # 日志级别
  folder: ./storage/log/    # 日志文件
  file: app_rotate.log      # 保存的日志文件
  rotate_count: 10          # 最多日志文件个数
  rotate_size: 1048576      # 每个日志大小（byte）
  rotate_time: "1d"         # 切割时间
  max_age: "90d"            # 文件保存时间
  date_format: "%Y-%m-%d"   # 文件后缀格式

# 按容量切割（缺少按日切割功能）
rolling:
  level: info             # 日志级别
  folder: ./storage/log/  # 日志文件
  file: app_rolling.log   # 保存的日志文件
  maxSize: 1              # 文件大小（mb） 0为不限制
  maxAge: 0               # 保留旧日志文件的最大天数 0为不限制
  maxBackups: 0           # 保留旧日志文件的最大数量 0为不限制

# 阿里云的 SLS 日志
aliyun_sls:
  level: info
  formatter: json
```
*/

// LogKey 定义字符串凭证
const LogKey = "gob:log"

type LogLevel uint32

const (
	// UnknownLevel 表示未知的日志级别
	UnknownLevel LogLevel = iota
	// PanicLevel level, panic 表示会导致整个程序出现崩溃的日志信息
	PanicLevel
	// FatalLevel level. fatal 表示会导致当前这个请求出现提前终止的错误信息
	FatalLevel
	// ErrorLevel level. error 表示出现错误，但是不一定影响后续请求逻辑的错误信息
	ErrorLevel
	// WarnLevel level. warn 表示出现错误，但是一定不影响后续请求逻辑的报警信息
	WarnLevel
	// InfoLevel level. info 表示正常的日志信息输出
	InfoLevel
	// DebugLevel level. debug 表示在调试状态下打印出来的日志信息
	DebugLevel
	// TraceLevel level. trace 表示最详细的信息，一般信息量比较大，可能包含调用堆栈等信息
	TraceLevel
)

// CtxFielder 定义了从context中获取信息的方法
type CtxFielder func(ctx context.Context) map[string]interface{}

// Formatter 定义了将日志信息组织成字符串的通用方法
type Formatter func(level LogLevel, t time.Time, msg string, fields map[string]interface{}) ([]byte, error)

// Log 定义了日志服务协议
type Log interface {
	// Panic 表示会导致整个程序出现崩溃的日志信息
	Panic(ctx context.Context, msg string, fields map[string]interface{})
	// Fatal 表示会导致当前这个请求出现提前终止的错误信息
	Fatal(ctx context.Context, msg string, fields map[string]interface{})
	// Error 表示出现错误，但是不一定影响后续请求逻辑的错误信息
	Error(ctx context.Context, msg string, fields map[string]interface{})
	// Warn 表示出现错误，但是一定不影响后续请求逻辑的报警信息
	Warn(ctx context.Context, msg string, fields map[string]interface{})
	// Info 表示正常的日志信息输出
	Info(ctx context.Context, msg string, fields map[string]interface{})
	// Debug 表示在调试状态下打印出来的日志信息
	Debug(ctx context.Context, msg string, fields map[string]interface{})
	// Trace 表示最详细的信息，一般信息量比较大，可能包含调用堆栈等信息
	Trace(ctx context.Context, msg string, fields map[string]interface{})
}
