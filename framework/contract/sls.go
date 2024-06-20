package contract

import "github.com/aliyun/aliyun-log-go-sdk/producer"

/*
## 服务介绍：
sls 提供对接阿里云日志服务的服务，可以用于日志的收集。
## 支持命令：无
## 支持配置：

配置文件为 config/[env]/app.yaml ，如下是一个配置示例：

```yaml
# 阿里云SLS服务
sls:
	endpoint: cn-shanghai.log.aliyuncs.com
	access_key_id: your_access_key_id
	access_key_secret: your_access_key_secret
	project: gob-test
	logstore: gob_test_logstore
```
*/

const SLSKey = "gob:sls"

type SLSService interface {
	GetSLSInstance() (*producer.Producer, error)
	GetProject() (string, error)
	GetLogstore() (string, error)
}
