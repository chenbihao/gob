package command

import (
	"context"
	"fmt"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/cobra"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/util"
	"github.com/pkg/errors"
	"github.com/sevlyar/go-daemon"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

/**
	命令介绍：
		web app 业务应用控制命令
	前置需求：
	支持命令：
		./gob app start		启动一个 app 服务
			--address=:8080 		指定端口
			--daemon=true 		守护模式（win下不支持）
		./gob app state 	获取启动的 app 的信息
		./gob app stop 		停止已经启动的 app 服务
		./gob app restart 	重新启动一个 app 服务
	支持配置：
		app.address 		地址格式需符合 http.Server 的 Addr 格式
		app.close_wait 		优雅关闭超时时间
	支持环境变量：
		ADDRESS				地址格式需符合 http.Server 的 Addr 格式
**/

var appAddress = ""   // app 启动地址
var appDaemon = false // app 守护模式

// initAppCommand 初始化app命令和其子命令
func initAppCommand() *cobra.Command {

	appStartCommand.Flags().StringVar(&appAddress, "address", ":8080", "设置app启动的地址，默认为:8080端口")
	appStartCommand.Flags().BoolVarP(&appDaemon, "daemon", "d", false, "以守护进程方式启动")

	appCommand.AddCommand(appStartCommand)
	appCommand.AddCommand(appRestartCommand)
	appCommand.AddCommand(appStateCommand)
	appCommand.AddCommand(appStopCommand)
	return appCommand
}

// AppCommand 是命令行参数第一级为app的命令，它没有实际功能，只是打印帮助文档
var appCommand = &cobra.Command{
	Use:   "app",
	Short: "业务应用控制命令",
	Long:  "业务应用控制命令，其包含业务启动，关闭，重启，查询等功能",
	RunE: func(c *cobra.Command, args []string) error {
		// 打印帮助文档
		c.Help()
		return nil
	},
}

// 启动AppServer, 这个函数会将当前goroutine阻塞
func startAppServe(c framework.Container, server *http.Server) error {
	// 这个goroutine是启动服务的goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	closeWait := 5
	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	if configService.IsExist("app.close_wait") {
		closeWait = configService.GetInt("app.close_wait")
	}
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Duration(closeWait)*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
		return err
	}
	return nil
}

// appStartCommand 启动一个Web服务
var appStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动一个app服务",
	RunE: func(c *cobra.Command, args []string) error {
		// 从Command中获取服务容器
		container := c.GetContainer()
		// 从服务容器中获取kernel的服务实例
		kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
		// 从kernel服务实例中获取引擎
		core := kernelService.HttpEngine()

		if appAddress == "" {
			envService := container.MustMake(contract.EnvKey).(contract.Env)
			if envService.Get("ADDRESS") != "" {
				appAddress = envService.Get("ADDRESS")
			} else {
				configService := container.MustMake(contract.ConfigKey).(contract.Config)
				if configService.IsExist("app.address") {
					appAddress = configService.GetString("app.address")
				} else {
					appAddress = ":8080"
				}
			}
		}
		// 创建一个Server服务
		server := &http.Server{
			Handler: core,
			Addr:    appAddress,
		}

		// 设置app的日志地址和进程id地址
		appService := container.MustMake(contract.AppKey).(contract.App)

		runtimeFolder := appService.RuntimeFolder()
		serverPidFile := filepath.Join(runtimeFolder, "app.pid")
		if err := util.CreateFolderIfNotExists(runtimeFolder); err != nil {
			return err
		}

		logFolder := appService.LogFolder()
		serverLogFile := filepath.Join(logFolder, "app.log")
		if err := util.CreateFolderIfNotExists(logFolder); err != nil {
			return err
		}
		currentFolder := util.GetExecDirectory()

		// daemon 模式
		if appDaemon {
			// win不支持 daemon 模式
			if util.IsWindows() {
				return errors.New("daemon: Non-POSIX OS is not supported")
			}
			// 创建一个Context
			cntxt := &daemon.Context{
				// 设置pid文件
				PidFileName: serverPidFile,
				PidFilePerm: 0664,
				// 设置日志文件
				LogFileName: serverLogFile,
				LogFilePerm: 0640,
				// 设置工作路径
				WorkDir: currentFolder,
				// 设置所有设置文件的mask，默认为750
				Umask: 027,
				// 子进程的参数，按照这个参数设置，子进程的命令为 ./gob app start --daemon=true
				Args: []string{"", "app", "start", "--daemon=true"},
			}
			// 启动子进程，d不为空表示当前是父进程，d为空表示当前是子进程
			// 这里可以把 Reborn 理解成 fork，当调用这个函数的时候，父进程会继续往下走，但是返回值 d 不为空，它的信息是子进程的进程号等信息。
			// 而子进程会重新运行对应的命令，再次进入到 Reborn 函数的时候，返回的 d 就为 nil
			child, err := cntxt.Reborn()
			if err != nil {
				return err
			}
			// 父进程直接打印启动成功信息，不做任何操作
			if child != nil {
				fmt.Println("app启动成功，pid:", child.Pid)
				fmt.Println("日志文件:", serverLogFile)
				return nil
			}
			defer cntxt.Release()
			// 子进程执行真正的app启动操作
			fmt.Println("daemon started")
			util.SetProcessTitle("gob app")
			if err := startAppServe(container, server); err != nil {
				fmt.Println(err)
			}
			return nil
		}

		// 非 daemon 模式，直接执行
		content := strconv.Itoa(os.Getpid())
		fmt.Println("[PID]", content)
		err := os.WriteFile(serverPidFile, []byte(content), 0644)
		if err != nil {
			return err
		}
		util.SetProcessTitle("gob app")

		fmt.Println("app serve url", appAddress)
		if err := startAppServe(container, server); err != nil {
			fmt.Println(err)
		}
		return nil
	},
}

// 重新启动一个app服务
var appRestartCommand = &cobra.Command{
	Use:   "restart",
	Short: "重新启动一个app服务",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		// GetPid
		runtimeFolder := appService.RuntimeFolder()
		serverPidFile := filepath.Join(runtimeFolder, "app.pid")
		if err := util.CreateFolderFileIfNotExists(runtimeFolder, serverPidFile); err != nil {
			return err
		}
		content, err := os.ReadFile(serverPidFile)
		if err != nil {
			return err
		}

		if len(content) != 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				return err
			}
			if util.CheckProcessExist(pid) {
				// 杀死进程
				if err := util.KillProcess(pid, syscall.SIGTERM); err != nil {
					return err
				}
				// 获取closeWait
				closeWait := 5
				configService := container.MustMake(contract.ConfigKey).(contract.Config)
				if configService.IsExist("app.close_wait") {
					closeWait = configService.GetInt("app.close_wait")
				}
				// 确认进程已经关闭,每秒检测一次， 最多检测 closeWait * 2秒
				for i := 0; i < closeWait*2; i++ {
					if !util.CheckProcessExist(pid) {
						break
					}
					time.Sleep(1 * time.Second)
				}
				// 如果进程等待了 2*closeWait 之后还没结束，返回错误，不进行后续的操作
				if util.CheckProcessExist(pid) {
					fmt.Println("结束进程失败:"+strconv.Itoa(pid), "请查看原因")
					return errors.New("结束进程失败")
				}
				//  清空 PID 文件
				if err := os.WriteFile(serverPidFile, []byte{}, 0644); err != nil {
					return err
				}
				fmt.Println("结束进程成功:" + strconv.Itoa(pid))
			}
		}

		// todo 这里 win 下后台运行未实现
		if util.IsNotWindows() {
			appDaemon = true
		}
		return appStartCommand.RunE(c, args)
	},
}

// 停止一个已经启动的app服务
var appStopCommand = &cobra.Command{
	Use:   "stop",
	Short: "停止一个已经启动的app服务",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		// GetPid
		serverPidFile := filepath.Join(appService.RuntimeFolder(), "app.pid")
		content, err := os.ReadFile(serverPidFile)
		if err != nil {
			return err
		}
		if len(content) != 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				return err
			}
			// 发送SIGTERM命令
			if err := util.KillProcess(pid, syscall.SIGTERM); err != nil {
				return err
			}
			if err := os.WriteFile(serverPidFile, []byte{}, 0644); err != nil {
				return err
			}
			fmt.Println("停止进程:", pid)
		}
		return nil
	},
}

// 获取启动的app的pid
var appStateCommand = &cobra.Command{
	Use:   "state",
	Short: "获取启动的app的pid",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		// 获取pid
		serverPidFile := filepath.Join(appService.RuntimeFolder(), "app.pid")
		content, err := os.ReadFile(serverPidFile)
		if err != nil {
			return err
		}
		if len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				return err
			}
			if util.CheckProcessExist(pid) {
				fmt.Println("app服务已经启动, pid:", pid)
				return nil
			}
		}
		fmt.Println("没有app服务存在")
		return nil
	},
}
