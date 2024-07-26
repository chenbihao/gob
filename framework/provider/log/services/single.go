package services

import (
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/util"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

type SingleLogService struct {
	SubLogService

	fd *os.File
}

// NewSingleLogService 实例化 SingleLog
func NewSingleLogService(parent *LogService) (interface{}, error) {

	log := &SingleLogService{}
	log.level = parent.parentLevel
	log.formatter = parent.parentFormatter
	log.folder = parent.parentFolder
	log.file = "app.log"

	config := parent.c.MustMake(contract.ConfigKey).(contract.Config)

	if exist, value := config.GetStringIfExist("log.single.level"); exist {
		log.level = GetLogLevel(value)
	}
	if exist, value := config.GetStringIfExist("log.single.formatter"); exist {
		log.formatter = GetLogFormatter(value)
	}
	if exist, value := config.GetStringIfExist("log.single.folder"); exist {
		log.folder = value
		_ = util.CreateFolderIfNotExists(log.folder)
	}
	if exist, value := config.GetStringIfExist("log.single.file"); exist {
		log.file = value
	}

	fd, err := os.OpenFile(filepath.Join(log.folder, log.file), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, errors.Wrap(err, "open log file err")
	}
	log.output = fd
	return log, nil
}
