package env

// 实现具体的服务实例 service.go

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/chenbihao/gob/framework/contract"
	"io"
	"os"
	"path"
	"strings"
)

// GobEnv 是 Env 的具体实现
type GobEnv struct {
	folder string            // 代表.env所在的目录
	maps   map[string]string // 保存所有的环境变量
}

// NewGobEnv 有一个参数，.env文件所在的目录
// example: NewGobEnv("/envfolder/") 会读取文件: /envfolder/.env
// .env的文件格式 FOO_ENV=BAR
func NewGobEnvService(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("NewGobEnv param error")
	}

	// 读取folder文件
	folder := params[0].(string)

	// 实例化
	gobEnv := &GobEnv{
		folder: folder,
		// 实例化环境变量，APP_ENV默认设置为开发环境
		maps: map[string]string{"APP_ENV": contract.EnvDev},
	}

	// 解析 folder/.env文件
	filePath := path.Join(folder, ".env")
	// 读取.env文件, 不管任意失败，都不影响后续

	// 打开文件.env
	file, err := os.Open(filePath)
	if err == nil {
		defer file.Close()

		// 读取文件
		br := bufio.NewReader(file)
		for {
			// 按照行进行读取
			line, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			// 按照等号解析
			s := bytes.SplitN(line, []byte{'='}, 2)
			// 如果不符合规范，则过滤
			if len(s) < 2 {
				continue
			}
			// 保存map
			key := string(s[0])
			val := string(s[1])
			gobEnv.maps[key] = val
		}
	}

	// 获取当前程序的环境变量，并且覆盖.env文件下的变量
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) < 2 {
			continue
		}
		gobEnv.maps[pair[0]] = pair[1]
	}

	// 返回实例
	return gobEnv, nil
}

// AppEnv 获取表示当前APP环境的变量APP_ENV
func (en *GobEnv) AppEnv() string {
	return en.Get("APP_ENV")
}

// IsExist 判断一个环境变量是否有被设置
func (en *GobEnv) IsExist(key string) bool {
	_, ok := en.maps[key]
	return ok
}

// Get 获取某个环境变量，如果没有设置，返回""
func (en *GobEnv) Get(key string) string {
	if val, ok := en.maps[key]; ok {
		return val
	}
	return ""
}

// All 获取所有的环境变量，.env和运行环境变量融合后结果
func (en *GobEnv) All() map[string]string {
	return en.maps
}
