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
	LogService
}

type SlsWriter struct {
	producer.Producer
	io.Writer
	c framework.Container
}

// NewSlsLog params sequence: level, ctxFielder, Formatter, map[string]interface(folder/file) aliyun sls log
func NewSlsLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	log := &SlsLog{}
	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)
	slsWriter, err := NewSlsWriter(c)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(slsWriter)
	log.c = c
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
