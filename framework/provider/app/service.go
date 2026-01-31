package app

// 实现具体的服务实例 service.go

import (
	"errors"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/util"
	"github.com/google/uuid"
)

// AppService 代表 gob 框架的 App 实现
type AppService struct {
	container  framework.Container // 服务容器
	appID      string              // 表示当前这个app的唯一id, 可以用于分布式锁等
	baseFolder string              // 基础路径
	// toolMode   bool                // 工具运行模式（通过 go install 安装至 $GOPATH/bin ）
	//argsMap   map[string]string // 参数加载(--key=value)
	//sysEnvMap map[string]string // 环境变量加载
	//configMap map[string]string // 配置加载
}

var _ contract.App = (*AppService)(nil)

// NewGobApp 初始化 AppService
func NewGobApp(params ...any) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}
	// 有两个参数，一个是容器，一个是 baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)

	uid, _ := uuid.NewV7()
	appID := uid.String()

	// toolMode := false
	// // 纯工具模式 ( 兼容 go install )
	// if os.Getenv("runMode") == "tool" || util.CheckBinaryFileInTheGOPATH() {
	// 	toolMode = true
	// }

	//configMap := map[string]string{}
	// gobApp := &AppService{baseFolder: baseFolder, container: container, appID: appID, configMap: configMap, toolMode: toolMode}
	gobApp := &AppService{baseFolder: baseFolder, container: container, appID: appID}
	//_ = gobApp.loadEnvMaps()
	//_ = gobApp.loadArgsMaps()
	return gobApp, nil

}

// AppID 表示当前这个app的唯一id, 可以用于分布式锁等
func (app *AppService) AppID() string {
	return app.appID
}

// Version 实现版本
func (app *AppService) Version() string {
	return framework.Version
}

// // IsToolMode 是否纯工具运行模式
// func (app *AppService) IsToolMode() bool {
// 	return app.toolMode
// }

// ---------------- 目录

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (app *AppService) BaseFolder() string {
	//if app.baseFolder != "" {
	//	return app.baseFolder
	//}
	//baseFolder := app.getConfigBySequence("base_folder", "BASE_FOLDER", "app.path.base_folder")
	//if baseFolder != "" {
	//	return baseFolder
	//}

	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}
