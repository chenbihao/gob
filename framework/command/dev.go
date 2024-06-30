package command

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/cobra"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/util"

	"github.com/fsnotify/fsnotify"
)

/*
## 命令介绍：
dev 调试工具，通过反向代理来管理前后端
## 前置需求：
前后端一体、cmd_build、config、app
## 支持命令：
```sh
./gob dev 			// 显示帮助信息
./gob dev frontend 	// 调试前端
./gob dev backend  	// 调试后端
./gob dev all  		// 显示所有
```
## 支持配置：
`app.yaml` 支持配置：
```yaml
dev: # 调试模式
port: 8070 # 调试模式最终监听的端口，默认为 8070
  frontend: # 前端调试模式配置
	port: 8071 # 前端监听端口, 默认 8071
  backend: # 后端调试模式配置
	refresh_time: 3  # 调试模式后端更新时间，如果文件变更，等待3s才进行一次更新，能让频繁保存变更更为顺畅, 默认1s
	port: 8072 # 后端监听端口，默认 8072
	monitor_folder: "" # 监听文件夹地址，为空或者不填默认为 AppFolder
```
*/

// 用于生成文档定位说明
const DevCommandKey = "调试模式命令"

// 初始化Dev命令
func initDevCommand() *cobra.Command {
	devCommand.AddCommand(devBackendCommand)
	devCommand.AddCommand(devFrontendCommand)
	devCommand.AddCommand(devAllCommand)
	return devCommand
}

// devCommand 为调试模式的一级命令
var devCommand = &cobra.Command{
	Use:   "dev",
	Short: "调试模式",
	RunE: func(c *cobra.Command, args []string) error {
		c.Help()
		return nil
	},
}

// devBackendCommand 启动后端调试模式
var devBackendCommand = &cobra.Command{
	Use:   "backend",
	Short: "启动后端调试模式",
	RunE: func(c *cobra.Command, args []string) error {
		proxy := NewProxy(c.GetContainer())
		go proxy.monitorBackend()
		if err := proxy.startProxy(false, true); err != nil {
			return err
		}
		return nil
	},
}

// devFrontendCommand 启动前端调试模式
var devFrontendCommand = &cobra.Command{
	Use:   "frontend",
	Short: "启动前端调试模式",
	RunE: func(c *cobra.Command, args []string) error {
		// 启动前端服务
		proxy := NewProxy(c.GetContainer())
		return proxy.startProxy(true, false)
	},
}

var devAllCommand = &cobra.Command{
	Use:   "all",
	Short: "同时启动前端和后端调试",
	RunE: func(c *cobra.Command, args []string) error {
		proxy := NewProxy(c.GetContainer())
		go proxy.monitorBackend()
		if err := proxy.startProxy(true, true); err != nil {
			return err
		}
		return nil
	},
}

// =========================================

// devConfig 代表调试模式的配置信息
type devConfig struct {
	Port    string   // 调试模式最终监听的端口，默认为 8070
	Backend struct { // 后端调试模式配置
		RefreshTime   int    // 调试模式后端更新时间，如果文件变更，等待3s才进行一次更新，能让频繁保存变更更为顺畅, 默认1s
		Port          string // 后端监听端口， 默认 8072
		MonitorFolder string // 监听文件夹，默认为AppFolder
	}
	Frontend struct { // 前端调试模式配置
		Port string // 前端启动端口, 默认 8071
	}
}

// 初始化配置文件
func initDevConfig(c framework.Container) *devConfig {
	// 设置默认值
	config := &devConfig{
		Port: "8070",
		Backend: struct {
			RefreshTime   int
			Port          string
			MonitorFolder string
		}{
			1,
			"8072",
			"",
		},
		Frontend: struct {
			Port string
		}{
			"8071",
		},
	}
	// 容器中获取配置服务
	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	// 每个配置项进行检查
	if configService.IsExist("app.dev.port") {
		config.Port = configService.GetString("app.dev.port")
	}
	if configService.IsExist("app.dev.backend.refresh_time") {
		config.Backend.RefreshTime = configService.GetInt("app.dev.backend.refresh_time")
	}
	if configService.IsExist("app.dev.backend.port") {
		config.Backend.Port = configService.GetString("app.dev.backend.port")
	}
	// monitorFolder 默认使用目录服务的 AppFolder()
	monitorFolder := configService.GetString("app.dev.backend.monitor_folder")
	if monitorFolder == "" {
		appService := c.MustMake(contract.AppKey).(contract.App)
		config.Backend.MonitorFolder = appService.AppFolder()
	}

	if configService.IsExist("app.dev.frontend.port") {
		config.Frontend.Port = configService.GetString("app.dev.frontend.port")
	}
	return config
}

// Proxy 代表serve启动的服务器代理
type Proxy struct {
	devConfig   *devConfig // 配置文件
	backendPid  int        // 当前的 backend 服务的 pid
	frontendPid int        // 当前的 frontend 服务的 pid
}

// NewProxy 初始化一个Proxy
func NewProxy(c framework.Container) *Proxy {
	config := initDevConfig(c)
	return &Proxy{
		devConfig: config,
	}
}

// 重新启动一个proxy网关
func (p *Proxy) newProxyReverseProxy(frontend, backend *url.URL) *httputil.ReverseProxy {
	if p.frontendPid == 0 && p.backendPid == 0 {
		fmt.Println("前端和后端服务都不存在")
		return nil
	}

	// 后端服务存在
	if p.frontendPid == 0 && p.backendPid != 0 {
		return httputil.NewSingleHostReverseProxy(backend)
	}
	// 前端服务存在
	if p.backendPid == 0 && p.frontendPid != 0 {
		return httputil.NewSingleHostReverseProxy(frontend)
	}

	// 两个都有进程
	// 先创建一个后端服务的 directory
	director := func(req *http.Request) {
		if req.URL.Path == "/" || req.URL.Path == "/app.js" {
			req.URL.Scheme = frontend.Scheme
			req.URL.Host = frontend.Host
		} else {
			req.URL.Scheme = backend.Scheme
			req.URL.Host = backend.Host
		}
	}

	// 定义一个 NotFoundErr
	NotFoundErr := errors.New("response is 404, need to redirect")
	return &httputil.ReverseProxy{
		Director: director, // 先转发到后端服务
		ModifyResponse: func(response *http.Response) error {
			// 如果后端服务返回了404，我们返回 NotFoundErr 会进入到 errorHandler 中
			if response.StatusCode == 404 {
				return NotFoundErr
			}
			return nil
		},
		ErrorHandler: func(writer http.ResponseWriter, request *http.Request, err error) {
			// 判断 Error 是否为NotFoundError, 是的话则进行前端服务的转发，重新修改writer
			if errors.Is(err, NotFoundErr) {
				httputil.NewSingleHostReverseProxy(frontend).ServeHTTP(writer, request)
			}
		}}
}

// 启动proxy服务，并且根据参数启动前端服务或者后端服务
func (p *Proxy) startProxy(startFrontend, startBackend bool) (err error) {
	var backendURL, frontendURL *url.URL
	// 启动后端
	if startBackend {
		if err = p.firstBuildBackend(); err != nil {
			fmt.Println("第一次编译失败：", err.Error())
			return
		}
		if err = p.restartBackend(); err != nil {
			return
		}
	}
	// 启动前端
	if startFrontend {
		if err = p.restartFrontend(); err != nil {
			return
		}
	}

	if frontendURL, err = url.Parse(fmt.Sprintf("%s%s", "http://127.0.0.1:", p.devConfig.Frontend.Port)); err != nil {
		return
	}
	if backendURL, err = url.Parse(fmt.Sprintf("%s%s", "http://127.0.0.1:", p.devConfig.Backend.Port)); err != nil {
		return
	}

	// 设置反向代理
	proxyReverse := p.newProxyReverseProxy(frontendURL, backendURL)
	proxyServer := &http.Server{
		Addr:    "127.0.0.1:" + p.devConfig.Port,
		Handler: proxyReverse,
	}

	fmt.Println("代理服务启动:", "http://"+proxyServer.Addr)
	// 启动proxy服务
	if err = proxyServer.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
	return nil
}

// rebuildBackend 重新编译后端
func (p *Proxy) rebuildBackend() error {
	// 重新编译
	bin := os.Args[0]
	cmdBuild := exec.Command(bin, "build", "backend")
	cmdBuild.Stdout = os.Stdout
	cmdBuild.Stderr = os.Stderr
	if err := cmdBuild.Start(); err == nil {
		err = cmdBuild.Wait()
		if err != nil {
			return err
		}
	}
	return nil
}

// firstBuildBackend 第一次编译后端
func (p *Proxy) firstBuildBackend() error {
	// 重新编译
	cmdBuild := exec.Command("go", "run", ".", "build", "backend")
	cmdBuild.Stdout = os.Stdout
	cmdBuild.Stderr = os.Stderr
	if err := cmdBuild.Start(); err == nil {
		if err = cmdBuild.Wait(); err != nil {
			return err
		}
	}
	return nil
}

// restartBackend 启动后端服务
func (p *Proxy) restartBackend() error {
	// 杀死之前的进程
	if p.backendPid != 0 {
		util.KillProcess(p.backendPid)
		p.backendPid = 0
	}

	// 设置随机端口，真实后端的端口
	port := p.devConfig.Backend.Port
	gobAddress := fmt.Sprintf(":" + port)
	bin := os.Args[0]
	// 使用命令行启动后端进程
	cmd := exec.Command(bin, "app", "start", "--address="+gobAddress)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("启动后端服务: ", "http://127.0.0.1:"+port)
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}
	p.backendPid = cmd.Process.Pid
	fmt.Println("后端服务pid:", p.backendPid)
	return nil
}

// 启动前端服务
func (p *Proxy) restartFrontend() error {
	// 杀死之前的进程
	if p.frontendPid != 0 {
		util.KillProcess(p.frontendPid)
		p.frontendPid = 0
	}

	// 开启 npm run dev
	port := p.devConfig.Frontend.Port
	path, err := exec.LookPath("npm")
	if err != nil {
		return err
	}

	// 把 port 参数传递进npm脚本里
	cmd := exec.Command(path, "run", "dev", "--", "--port", port)
	cmd.Dir = frontendFolder
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("启动前端服务: ", "http://127.0.0.1:"+port)
	if err = cmd.Start(); err != nil {
		fmt.Println(err)
	}
	p.frontendPid = cmd.Process.Pid
	fmt.Println("前端服务pid:", p.frontendPid)
	return nil
}

// monitorBackend 监听应用文件
func (p *Proxy) monitorBackend() (err error) {
	// 监听
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}
	defer watcher.Close()

	// 开启监听目标文件夹
	appFolder := p.devConfig.Backend.MonitorFolder
	fmt.Println("监控文件夹：", appFolder)
	// 监听所有子目录，需要使用filepath.walk
	filepath.Walk(appFolder, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			return nil
		}
		// 如果是隐藏的目录比如 . 或者 .. 则不用进行监控
		if util.IsHiddenDirectory(path) {
			return nil
		}
		return watcher.Add(path)
	})

	// 开启计时时间机制
	refreshTime := p.devConfig.Backend.RefreshTime
	t := time.NewTimer(time.Duration(refreshTime) * time.Second)
	// 先停止计时器
	t.Stop()
	for {
		select {
		case <-t.C:
			// 计时器时间到了，代表之前有文件更新事件重置过计时器
			// 即有文件更新
			fmt.Println("...检测到文件更新，重启服务开始...")
			fmt.Println("...期间请不要发送任何请求...")
			if err := p.rebuildBackend(); err != nil {
				fmt.Println("重新编译失败：", err.Error())
			} else {
				if err := p.restartBackend(); err != nil {
					fmt.Println("重新启动失败：", err.Error())
				}
			}
			fmt.Println("...检测到文件更新，重启服务结束...")
			// 停止计时器
			t.Stop()
		case _, ok := <-watcher.Events:
			if !ok {
				continue
			}
			// 有文件更新事件，重置计时器
			t.Reset(time.Duration(refreshTime) * time.Second)
		case err, ok := <-watcher.Errors:
			if !ok {
				continue
			}
			// 如果有文件监听错误，则停止计时器
			fmt.Println("监听文件夹错误：", err.Error())
			t.Reset(time.Duration(refreshTime) * time.Second)
		}
	}
}
