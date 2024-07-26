package services

import (
	"fmt"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/util"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"path/filepath"
	"time"
)

// RotateLogService 代表会进行切割的日志文件存储
type RotateLogService struct {
	SubLogService
}

// NewRotateLogService 实例化 RotateLogService
func NewRotateLogService(parent *LogService) (interface{}, error) {

	log := &RotateLogService{}
	log.level = parent.parentLevel
	log.formatter = parent.parentFormatter
	log.folder = parent.parentFolder
	log.file = "app.log"

	config := parent.c.MustMake(contract.ConfigKey).(contract.Config)

	if exist, value := config.GetStringIfExist("log.rotate.level"); exist {
		log.level = GetLogLevel(value)
	}
	if exist, value := config.GetStringIfExist("log.rotate.formatter"); exist {
		log.formatter = GetLogFormatter(value)
	}
	if exist, value := config.GetStringIfExist("log.rotate.folder"); exist {
		log.folder = value
		_ = util.CreateFolderIfNotExists(log.folder)
	}
	if exist, value := config.GetStringIfExist("log.rotate.file"); exist {
		log.file = value
	}

	linkName := rotatelogs.WithLinkName(filepath.Join(log.folder, log.file))
	options := []rotatelogs.Option{linkName}

	// 从配置文件获取date_format信息
	dateFormat := "%Y%m%d"
	if exist, value := config.GetStringIfExist("log.rotate.date_format"); exist {
		dateFormat = value
	}
	// 从配置文件获取 rotate_count 信息
	if exist, value := config.GetIntIfExist("log.rotate.rotate_count"); exist {
		options = append(options, rotatelogs.WithRotationCount(uint(value)))
	}
	// 从配置文件获取 rotate_size 信息
	if exist, value := config.GetIntIfExist("log.rotate.rotate_size"); exist {
		options = append(options, rotatelogs.WithRotationSize(int64(value)))
	}
	// 从配置文件获取 max_age 信息
	if exist, value := config.GetStringIfExist("log.rotate.max_age"); exist {
		if maxAgeParse, err := time.ParseDuration(value); err == nil {
			options = append(options, rotatelogs.WithMaxAge(maxAgeParse))
		}
	}
	// 从配置文件获取rotate_time信息
	if exist, value := config.GetStringIfExist("log.rotate.rotate_time"); exist {
		if rotateTimeParse, err := time.ParseDuration(value); err == nil {
			options = append(options, rotatelogs.WithRotationTime(rotateTimeParse))
		}
	}

	w, err := rotatelogs.New(fmt.Sprintf("%s.%s", filepath.Join(log.folder, log.file), dateFormat), options...)
	if err != nil {
		return nil, errors.Wrap(err, "new rotate logs error")
	}
	log.output = w
	return log, nil
}
