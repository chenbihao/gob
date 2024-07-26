package services

import (
	"fmt"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"io"
	"time"
)

type SlsLog struct {
	SubLogService
}

type SlsWriter struct {
	c framework.Container
	producer.Producer
	io.Writer
}

// NewSlsLog
func NewSlsLog(parent *LogService) (interface{}, error) {
	log := &SlsLog{}
	log.level = parent.parentLevel
	log.formatter = parent.parentFormatter

	config := parent.c.MustMake(contract.ConfigKey).(contract.Config)

	if exist, value := config.GetStringIfExist("log.sls.level"); exist {
		log.level = GetLogLevel(value)
	}
	if exist, value := config.GetStringIfExist("log.sls.formatter"); exist {
		log.formatter = GetLogFormatter(value)
	}

	slsWriter, err := NewSlsWriter(parent.c)
	if err != nil {
		fmt.Println(err)
	}
	log.output = slsWriter
	return log, nil
}

func NewSlsWriter(c framework.Container) (*SlsWriter, error) {
	slsWriter := &SlsWriter{}
	slsWriter.c = c
	return slsWriter, nil
}

func (s *SlsWriter) Write(p []byte) (int, error) {
	slsService := s.c.MustMake(contract.SLSKey).(contract.SLSService)
	producerInstance, err := slsService.GetSLSInstance()
	if err != nil {
		panic(err)
	}

	project, err := slsService.GetProject()
	if err != nil {
		panic(err)
	}
	logstore, err := slsService.GetLogstore()
	if err != nil {
		panic(err)
	}

	ch := make(chan struct{})
	go func() {
		if string(p) == "\r\n" {
			ch <- struct{}{}
			return
		}
		logger := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{"content": string(p)})
		err := producerInstance.SendLog(project, logstore, "topic", "127.0.0.1", logger)
		if err != nil {
			fmt.Println(err)
		}
		ch <- struct{}{}
	}()

	if _, ok := <-ch; ok {
		fmt.Println("Send completion")
		//go func() {
		//	producerInstance.SafeClose()
		//}()
	}

	return len(p), nil
}
