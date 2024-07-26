package services

import (
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/util"
	"gopkg.in/natefinch/lumberjack.v2"
	"path/filepath"
)

// RollingLogService 代表会进行切割的日志文件存储
type RollingLogService struct {
	SubLogService
}

// NewRotateLogService 实例化 RotateLogService
func NewRollingLogService(parent *LogService) (interface{}, error) {

	log := &RollingLogService{}
	log.level = parent.parentLevel
	log.formatter = parent.parentFormatter
	log.folder = parent.parentFolder
	log.file = "app.log"

	config := parent.c.MustMake(contract.ConfigKey).(contract.Config)

	if exist, value := config.GetStringIfExist("log.rolling.level"); exist {
		log.level = GetLogLevel(value)
	}
	if exist, value := config.GetStringIfExist("log.rolling.formatter"); exist {
		log.formatter = GetLogFormatter(value)
	}
	if exist, value := config.GetStringIfExist("log.rolling.folder"); exist {
		log.folder = value
		_ = util.CreateFolderIfNotExists(log.folder)
	}
	if exist, value := config.GetStringIfExist("log.rolling.file"); exist {
		log.file = value
	}
	flieName := filepath.Join(log.folder, log.file)

	maxSize := 0
	maxAge := 0
	maxBackups := 0

	if exist, value := config.GetIntIfExist("log.rolling.maxSize"); exist {
		maxSize = value
	}
	if exist, value := config.GetIntIfExist("log.rolling.maxAge"); exist {
		maxAge = value
	}
	if exist, value := config.GetIntIfExist("log.rolling.maxBackups"); exist {
		maxBackups = value
	}

	output := &lumberjack.Logger{
		Filename:   flieName,
		MaxSize:    maxSize,    // megabytes
		MaxAge:     maxAge,     // days
		MaxBackups: maxBackups, //
		Compress:   false,
		LocalTime:  true,
	}
	log.output = output
	return log, nil
}
