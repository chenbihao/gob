package command

import (
	"context"
	"fmt"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/cobra"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/provider/ssh"
	"github.com/chenbihao/gob/framework/util"
	"github.com/pkg/errors"
	"github.com/pkg/sftp"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

/*
## 命令介绍：
部署命令
## 前置需求：
app
## 支持命令：
```sh
./gob deploy frontend`	// 部署前端
	-s --skip-build     	// 跳过前端构建
./gob deploy backend`	// 部署后端
./gob deploy all`		// 同时部署前后端
	-s --skip-build     	// 跳过前端构建
./gob deploy rollback`	// 部署回滚
```
## 支持配置：
`deploy.yaml` 支持配置：
```yaml
connections: # 要自动化部署的连接
  - ssh.web-key

remote_folder: "/home/demo/deploy/"  # 远端的部署文件夹

frontend: # 前端部署配置
  pre_action: # 部署前置命令
	- "pwd"
  post_action: # 部署后置命令
	- "pwd"

backend: # 后端部署配置
  goos: linux # 部署目标操作系统
  goarch: amd64 # 部署目标cpu架构
  pre_action: # 部署前置命令
	- "rm /home/demo/deploy/gob"
  post_action: # 部署后置命令
	- "chmod 777 /home/demo/deploy/gob"
	- "/home/demo/deploy/gob app restart"
```
ssh 支持配置：详见 contract/redis.go
*/

// 用于生成文档定位说明
const DeployCommandKey = "部署命令"

var deploySkipBuild = false // 跳过编译环节

// initDeployCommand 为自动化部署的命令
func initDeployCommand() *cobra.Command {
	deployFrontendCommand.Flags().BoolVarP(&deploySkipBuild, "skip-build", "s", false, "跳过构建环节")
	deployAllCommand.Flags().BoolVarP(&deploySkipBuild, "skip-build", "s", false, "跳过构建环节")
	deployCommand.AddCommand(deployFrontendCommand)
	deployCommand.AddCommand(deployBackendCommand)
	deployCommand.AddCommand(deployAllCommand)
	deployCommand.AddCommand(deployRollbackCommand)
	return deployCommand
}

// deployCommand 一级命令，显示帮助信息
var deployCommand = &cobra.Command{
	Use:   "deploy",
	Short: "部署相关命令",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			c.Help()
		}
		return nil
	},
}

// deployFrontendCommand 部署前端
var deployFrontendCommand = &cobra.Command{
	Use:   "frontend",
	Short: "部署前端",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		// 创建部署文件夹
		deployFolder, err := createDeployFolder(container)
		if err != nil {
			return err
		}
		// 编译前端到部署文件夹
		if err := deployBuildFrontend(c, deployFolder); err != nil {
			return err
		}
		// 上传部署文件夹并执行对应的shell
		return deployUploadAction(deployFolder, container, "frontend")
	},
}

// deployBackendCommand 部署后端
var deployBackendCommand = &cobra.Command{
	Use:   "backend",
	Short: "部署后端",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		// 创建部署文件夹
		deployFolder, err := createDeployFolder(container)
		if err != nil {
			return err
		}
		// 编译后端到部署文件夹
		if err := deployBuildBackend(c, deployFolder); err != nil {
			return err
		}
		// 上传部署文件夹并执行对应的shell
		return deployUploadAction(deployFolder, container, "backend")
	},
}

var deployAllCommand = &cobra.Command{
	Use:   "all",
	Short: "全部部署",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		deployFolder, err := createDeployFolder(container)
		if err != nil {
			return err
		}
		// 编译前端
		if err := deployBuildFrontend(c, deployFolder); err != nil {
			return err
		}
		// 编译后端
		if err := deployBuildBackend(c, deployFolder); err != nil {
			return err
		}
		// 上传前端+后端，并执行对应的shell
		if err := deployUploadAction(deployFolder, container, "all"); err != nil {
			return err
		}
		return nil
	},
}

// deployRollbackCommand 部署回滚
var deployRollbackCommand = &cobra.Command{
	Use:   "rollback",
	Short: "部署回滚",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()

		if len(args) != 2 {
			return errors.New("参数错误,请按照参数进行回滚 ./gob deploy rollback [version] [frontend/backend/all]")
		}

		version := args[0]
		end := args[1]

		// 获取版本信息
		appService := container.MustMake(contract.AppKey).(contract.App)
		deployFolder := filepath.Join(appService.DeployFolder(), version)

		// 上传部署文件夹并执行对应的shell
		return deployUploadAction(deployFolder, container, end)
	},
}

func deployBuildBackend(c *cobra.Command, deployFolder string) error {
	container := c.GetContainer()
	configService := container.MustMake(contract.ConfigKey).(contract.Config)
	appService := container.MustMake(contract.AppKey).(contract.App)
	envService := container.MustMake(contract.EnvKey).(contract.Env)
	logger := container.MustMake(contract.LogKey).(contract.Log)

	env := envService.AppEnv()

	binFile := "gob"

	// 编译后端
	path, err := exec.LookPath("go")
	if err != nil {
		log.Fatalln("gob go: 请在Path路径中先安装go")
	}

	deployBinFile := filepath.Join(deployFolder, binFile)
	cmd := exec.Command(path, "build", "-o", deployBinFile, "./")
	cmd.Env = os.Environ()
	// 设置GOOS和GOARCH
	if configService.GetString("deploy.backend.goos") != "" {
		cmd.Env = append(cmd.Env, "GOOS="+configService.GetString("deploy.backend.goos"))
	}
	if configService.GetString("deploy.backend.goarch") != "" {
		cmd.Env = append(cmd.Env, "GOARCH="+configService.GetString("deploy.backend.goarch"))
	}

	// 执行命令
	ctx := context.Background()
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error(ctx, "go build err", map[string]interface{}{
			"err": err,
			"out": string(out),
		})
		return err
	}
	logger.Info(ctx, "后端编译成功", nil)

	// 复制.env文件
	if util.Exists(filepath.Join(appService.BaseFolder(), ".env")) {
		if err := util.CopyFile(filepath.Join(appService.BaseFolder(), ".env"), filepath.Join(deployFolder, ".env")); err != nil {
			return err
		}
	}

	// 复制config文件
	deployConfigFolder := filepath.Join(deployFolder, "config", env)
	if !util.Exists(deployConfigFolder) {
		if err := os.MkdirAll(deployConfigFolder, os.ModePerm); err != nil {
			return err
		}
	}
	if err := util.CopyFolder(filepath.Join(appService.ConfigFolder(), env), deployConfigFolder); err != nil {
		return err
	}

	logger.Info(ctx, "build local ok", nil)
	return nil
}

// 上传部署文件夹，并且执行对应的前置和后置的shell
func deployUploadAction(deployFolder string, container framework.Container, end string) error {
	configService := container.MustMake(contract.ConfigKey).(contract.Config)
	sshService := container.MustMake(contract.SSHKey).(contract.SSHService)
	logger := container.MustMake(contract.LogKey).(contract.Log)

	// 遍历所有deploy的服务器
	deployNodes := configService.GetStringSlice("deploy.connections")
	if len(deployNodes) == 0 {
		return errors.New("deploy connections len is zero")
	}
	remoteFolder := configService.GetString("deploy.remote_folder")
	if remoteFolder == "" {
		return errors.New("remote folder is empty")
	}

	preActions := make([]string, 0, 1)
	postActions := make([]string, 0, 1)

	if end == "frontend" || end == "all" {
		preActions = append(preActions, configService.GetStringSlice("deploy.frontend.pre_action")...)
		postActions = append(postActions, configService.GetStringSlice("deploy.frontend.post_action")...)
	}
	if end == "backend" || end == "all" {
		preActions = append(preActions, configService.GetStringSlice("deploy.backend.pre_action")...)
		postActions = append(postActions, configService.GetStringSlice("deploy.backend.post_action")...)
	}

	// 对每个远端服务
	for _, node := range deployNodes {
		sshClient, err := sshService.GetClient(ssh.WithConfigPath(node))
		if err != nil {
			return err
		}
		client, err := sftp.NewClient(sshClient)
		if err != nil {
			return err
		}

		// 执行所有的前置命令
		for _, action := range preActions {
			// 创建session
			session, err := sshClient.NewSession()
			if err != nil {
				return err
			}
			logger.Info(context.Background(), "execute pre action start", map[string]interface{}{
				"cmd":        action,
				"connection": node,
			})
			// 执行命令，并且等待返回
			bts, err := session.CombinedOutput(action)
			if err != nil && !strings.HasPrefix(action, "rm") {
				session.Close()
				return err
			}
			session.Close()
			// 执行前置命令成功
			logger.Info(context.Background(), "execute pre action", map[string]interface{}{
				"cmd":        action,
				"connection": node,
				"out":        strings.ReplaceAll(string(bts), "\n", ""),
			})
		}

		if err := uploadFolderToSFTP(container, deployFolder, remoteFolder, client); err != nil {
			logger.Info(context.Background(), "upload folder failed", map[string]interface{}{
				"err": err,
			})
			return err
		}
		logger.Info(context.Background(), "upload folder success", nil)

		for _, action := range postActions {
			session, err := sshClient.NewSession()
			if err != nil {
				return err
			}
			logger.Info(context.Background(), "execute post action start", map[string]interface{}{
				"cmd":        action,
				"connection": node,
			})
			bts, err := session.CombinedOutput(action)
			if err != nil {
				logger.Info(context.Background(), "execute post action err", map[string]interface{}{
					"cmd":        action,
					"connection": node,
					"out":        strings.ReplaceAll(string(bts), "\n", ""),
				})
				session.Close()
				return err
			}
			logger.Info(context.Background(), "execute post action finish", map[string]interface{}{
				"cmd":        action,
				"connection": node,
				"out":        strings.ReplaceAll(string(bts), "\n", ""),
			})
			session.Close()
		}
	}
	return nil
}

// 上传部署文件夹
func uploadFolderToSFTP(container framework.Container, localFolder, remoteFolder string, client *sftp.Client) error {
	logger := container.MustMake(contract.LogKey).(contract.Log)
	config := container.MustMake(contract.ConfigKey).(contract.Config)

	// 遍历本地文件
	return filepath.Walk(localFolder, func(path string, info os.FileInfo, err error) error {
		// 获取除了folder前缀的后续文件名称
		relPath := strings.Replace(path, localFolder, "", 1)
		if relPath == "" {
			return nil
		}

		// 这里做一道转换
		remoteFile := filepath.Join(remoteFolder, relPath)
		if config.GetString("deploy.backend.goos") == "linux" {
			remoteFile = filepath.ToSlash(remoteFile)
		}

		// 如果是遍历到了一个目录
		if info.IsDir() {
			logger.Info(context.Background(), "mkdir: "+remoteFile, nil)
			// 创建这个目录
			return client.MkdirAll(remoteFile)
		}

		// 打开本地的文件
		rf, err := os.Open(filepath.Join(localFolder, relPath))
		if err != nil {
			return errors.New("read file " + filepath.Join(localFolder, relPath) + " error:" + err.Error())
		}
		defer rf.Close()
		// 检查文件大小
		rfStat, err := rf.Stat()
		if err != nil {
			return err
		}

		// 打开/创建远端文件
		f, err := client.Create(remoteFile)
		if err != nil {
			return errors.New("create file " + remoteFile + " error:" + err.Error())
		}
		defer f.Close()

		// 大于2M的文件显示进度
		if rfStat.Size() > 2*1024*1024 {
			logger.Info(context.Background(), "upload local file: "+filepath.Join(localFolder, relPath)+
				" to remote file: "+remoteFile+" start", nil)
			// 开启一个goroutine来不断计算进度
			go func(localFile, remoteFile string) {
				// 每1s计算一次
				ticker := time.NewTicker(1 * time.Second)
				for range ticker.C {
					// 获取远端文件信息
					remoteFileInfo, err := client.Stat(remoteFile)
					if err != nil {
						logger.Error(context.Background(), "stat error", map[string]interface{}{
							"err":         err,
							"remote_file": remoteFile,
						})
						continue
					}
					// 如果远端文件大小等于本地文件大小，说明已经结束了
					size := remoteFileInfo.Size()
					if size >= rfStat.Size() {
						break
					}
					// 计算进度并且打印进度
					percent := int(size * 100 / rfStat.Size())
					logger.Info(context.Background(), "upload local file: "+filepath.Join(localFolder, relPath)+
						" to remote file: "+remoteFile+fmt.Sprintf(" %v%% %v/%v", percent, size, rfStat.Size()), nil)
				}
			}(filepath.Join(localFolder, relPath), remoteFile)
		}

		// 将本地文件并发读取到远端文件
		if _, err := f.ReadFromWithConcurrency(rf, 10); err != nil {
			return errors.New("Write file " + remoteFile + " error:" + err.Error())
		}
		// 记录成功信息
		logger.Info(context.Background(), "upload local file: "+filepath.Join(localFolder, relPath)+
			" to remote file: "+remoteFile+" finish", nil)
		return nil
	})
}

func deployBuildFrontend(c *cobra.Command, deployFolder string) error {
	container := c.GetContainer()
	appService := container.MustMake(contract.AppKey).(contract.App)

	if !deploySkipBuild {
		// 编译前端
		if err := buildFrontendCommand.RunE(c, []string{}); err != nil {
			return err
		}
	}

	// 复制前端文件到deploy文件夹
	frontendFolder := filepath.Join(deployFolder, "dist")
	if err := os.Mkdir(frontendFolder, os.ModePerm); err != nil {
		return err
	}
	buildFolder := filepath.Join(appService.BaseFolder(), "gob_frontend/dist")
	if err := util.CopyFolder(buildFolder, frontendFolder); err != nil {
		return err
	}
	return nil
}

// 创建部署的folder
func createDeployFolder(c framework.Container) (string, error) {
	appService := c.MustMake(contract.AppKey).(contract.App)
	deployFolder := appService.DeployFolder()

	deployVersion := time.Now().Format("20060102150405")
	versionFolder := filepath.Join(deployFolder, deployVersion)
	if !util.Exists(versionFolder) {
		return versionFolder, os.Mkdir(versionFolder, os.ModePerm)
	}
	return versionFolder, nil
}
