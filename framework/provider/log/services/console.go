package services

import (
	"github.com/chenbihao/gob/framework/contract"
	"os"
)

// ConsoleLogService 代表控制台输出
type ConsoleLogService struct {
	SubLogService
}

// NewConsoleLogService 实例化 ConsoleLog
func NewConsoleLogService(parent *LogService) (interface{}, error) {

	log := &ConsoleLogService{}
	log.level = parent.parentLevel
	log.formatter = parent.parentFormatter

	config := parent.c.MustMake(contract.ConfigKey).(contract.Config)

	if exist, value := config.GetStringIfExist("log.console.level"); exist {
		log.level = GetLogLevel(value)
	}
	if exist, value := config.GetStringIfExist("log.console.formatter"); exist {
		log.formatter = GetLogFormatter(value)
	}

	// 最重要的将内容输出到控制台
	log.output = os.Stdout
	return log, nil
}
