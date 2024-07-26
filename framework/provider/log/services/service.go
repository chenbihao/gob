package services

import (
	"context"
	"github.com/chenbihao/gob/framework/provider/log/formatter"
	"github.com/chenbihao/gob/framework/util"
	"io"
	pkgLog "log"
	"strings"
	"time"

	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
)

// LogService 的通用实例
type LogService struct {
	c               framework.Container // 容器
	ctxFielder      contract.CtxFielder // ctx 获取上下文字段
	parentLevel     contract.LogLevel   // 父日志级别
	parentFormatter contract.Formatter  // 父日志格式化方法
	parentFolder    string              // 父日志存放目录
	subLogService   []SubLogService     // 子日志服务
}

type SubLogService struct {
	level     contract.LogLevel  // 日志级别
	formatter contract.Formatter // 日志格式化方法
	folder    string             // 日志存放目录
	file      string             // 日志文件名
	output    io.Writer          // 输出
}

// NewLogService 实例化 LogService
func NewLogService(params ...interface{}) (interface{}, error) {

	c := params[0].(framework.Container)
	drivers := params[1].([]string)
	ctxFielder := params[2].(contract.CtxFielder)

	app := c.MustMake(contract.AppKey).(contract.App)
	config := c.MustMake(contract.ConfigKey).(contract.Config)

	parentLevel := contract.InfoLevel
	if exist, value := config.GetStringIfExist("log.level"); exist {
		parentLevel = GetLogLevel(value)
	}
	parentFormatter := formatter.TextFormatter
	if exist, value := config.GetStringIfExist("log.formatter"); exist {
		parentFormatter = GetLogFormatter(value)
	}
	parentFolder := app.LogFolder()
	if exist, value := config.GetStringIfExist("log.folder"); exist {
		parentFolder = value
	}
	_ = util.CreateFolderIfNotExists(parentFolder)

	log := &LogService{}
	log.c = c
	log.ctxFielder = ctxFielder
	log.parentLevel = parentLevel
	log.parentFormatter = parentFormatter
	log.parentFolder = parentFolder

	// 读取配置 按照配置列表来载入相关的日志
	if len(drivers) == 0 {
		drivers = append(drivers, "console")
	}
	for _, driver := range drivers {
		switch driver {
		case "console":
			consoleLog, _ := NewConsoleLogService(log)
			log.subLogService = append(log.subLogService, consoleLog.(*ConsoleLogService).SubLogService)
		case "single":
			singleLog, _ := NewSingleLogService(log)
			log.subLogService = append(log.subLogService, singleLog.(*SingleLogService).SubLogService)
		case "rotate":
			rotateLog, _ := NewRotateLogService(log)
			log.subLogService = append(log.subLogService, rotateLog.(*RotateLogService).SubLogService)
		case "rolling":
			rollingLog, _ := NewRollingLogService(log)
			log.subLogService = append(log.subLogService, rollingLog.(*RollingLogService).SubLogService)
		case "aliyun_sls":
			slsLog, _ := NewSlsLog(log)
			log.subLogService = append(log.subLogService, slsLog.(*SlsLog).SubLogService)
		}
	}
	return log, nil
}

// IsLevelEnable 判断这个级别是否可以打印
func (log *SubLogService) IsLevelEnable(level contract.LogLevel) bool {
	return level <= log.level
}

// logf 为打印日志的核心函数
func (log *LogService) logf(ctx context.Context, level contract.LogLevel, msg string, fields map[string]interface{}) error {

	// 使用 ctxFielder 获取 context 中的信息
	fs := fields
	if fs == nil {
		fs = make(map[string]interface{})
	}
	if log.ctxFielder != nil {
		t := log.ctxFielder(ctx)
		for k, v := range t {
			fs[k] = v
		}
	}

	// 如果绑定了 trace 服务，获取trace信息
	if log.c != nil && log.c.IsBind(contract.TraceKey) {
		tracer := log.c.MustMake(contract.TraceKey).(contract.Trace)
		tc := tracer.GetTrace(ctx)
		if tc != nil {
			maps := tracer.ToMap(tc)
			for k, v := range maps {
				fs[k] = v
			}
		}
	}

	// 遍历注册的子服务
	for i := range log.subLogService {
		// 打印相应等级日志
		subLog := &log.subLogService[i]
		if !subLog.IsLevelEnable(level) {
			continue
		}
		// 将日志信息按照 formatter 序列化为字符串
		ct, err := subLog.formatter(level, time.Now(), msg, fs)
		if err != nil {
			return err
		}
		// 如果是panic级别，则使用log进行panic
		if level == contract.PanicLevel {
			pkgLog.Panicln(string(ct))
			continue
		}
		// 通过 output 进行输出
		_, _ = subLog.output.Write(ct)
		_, _ = subLog.output.Write([]byte("\r\n"))
	}

	return nil
}

// Panic 输出panic的日志信息
func (log *LogService) Panic(ctx context.Context, msg string, fields map[string]interface{}) {
	_ = log.logf(ctx, contract.PanicLevel, msg, fields)
}

// Fatal will add fatal record which contains msg and fields
func (log *LogService) Fatal(ctx context.Context, msg string, fields map[string]interface{}) {
	_ = log.logf(ctx, contract.FatalLevel, msg, fields)
}

// Error will add error record which contains msg and fields
func (log *LogService) Error(ctx context.Context, msg string, fields map[string]interface{}) {
	_ = log.logf(ctx, contract.ErrorLevel, msg, fields)
}

// Warn will add warn record which contains msg and fields
func (log *LogService) Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	_ = log.logf(ctx, contract.WarnLevel, msg, fields)
}

// Info 会打印出普通的日志信息
func (log *LogService) Info(ctx context.Context, msg string, fields map[string]interface{}) {
	_ = log.logf(ctx, contract.InfoLevel, msg, fields)
}

// Debug will add debug record which contains msg and fields
func (log *LogService) Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	_ = log.logf(ctx, contract.DebugLevel, msg, fields)
}

// Trace will add trace info which contains msg and fields
func (log *LogService) Trace(ctx context.Context, msg string, fields map[string]interface{}) {
	_ = log.logf(ctx, contract.TraceLevel, msg, fields)
}

func GetLogLevel(str string) contract.LogLevel {
	switch strings.ToLower(str) {
	case "panic":
		return contract.PanicLevel
	case "fatal":
		return contract.FatalLevel
	case "error":
		return contract.ErrorLevel
	case "warn":
		return contract.WarnLevel
	case "info":
		return contract.InfoLevel
	case "debug":
		return contract.DebugLevel
	case "trace":
		return contract.TraceLevel
	default:
		return contract.UnknownLevel
	}
}

func GetLogFormatter(str string) contract.Formatter {
	switch strings.ToLower(str) {
	case "json":
		return formatter.JsonFormatter
	case "text":
		return formatter.TextFormatter
	default:
		return formatter.TextFormatter
	}
}
