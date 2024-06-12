//go:build windows

package distributed

import (
	"errors"
	"github.com/chenbihao/gob/framework"
	"github.com/chenbihao/gob/framework/contract"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// LocalDistributedService
type LocalDistributedService struct {
	container framework.Container // 服务容器
}

var _ contract.Distributed = (*LocalDistributedService)(nil)

// NewLocalDistributedService 初始化本地分布式服务
func NewLocalDistributedService(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("param error")
	}

	container := params[0].(framework.Container)
	return &LocalDistributedService{container: container}, nil
}

// Select 为分布式选择器   （这里win不支持syscall.Flock，作废，换兼容性更好的）
func (s LocalDistributedService) Select(serviceName string, appID string, holdTime time.Duration) (selectAppID string, err error) {
	appService := s.container.MustMake(contract.AppKey).(contract.App)
	runtimeFolder := appService.RuntimeFolder()
	lockFile := filepath.Join(runtimeFolder, "disribute_"+serviceName+".txt")

	// 打开文件锁
	lock, err := os.OpenFile(lockFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}

	// win 下使用文件锁
	h, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return
	}
	defer syscall.FreeLibrary(h)

	addr, err := syscall.GetProcAddress(h, "LockFile")
	if err != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(addr, 5, lock.Fd(), 0, 0, 0, 1, 0)
	if 0 == int(r0) {
		// 加锁失败，只读的形式读取占用中的appid
		selectAppIDByt, readErr := io.ReadAll(lock)
		if readErr != nil {
			if strings.Contains(readErr.Error(), "another process has locked a portion of the file.") {
				return "ReadFileFailed", nil
			}
			return "", readErr
		}
		return string(selectAppIDByt), err
	}

	// 在一段时间内，选举有效，其他节点在这段时间不能再进行抢占
	go func() {
		defer func() {
			// 释放文件锁
			addr, err = syscall.GetProcAddress(h, "UnlockFile")
			if err != nil {
				return
			}
			syscall.Syscall6(addr, 5, lock.Fd(), 0, 0, 0, 1, 0)

			// 释放文件
			lock.Close()
			// 删除文件锁对应的文件
			os.Remove(lockFile)
		}()
		// 创建选举结果有效的计时器
		timer := time.NewTimer(holdTime)
		// 等待计时器结束
		<-timer.C
	}()

	// 这里已经是抢占到了，将抢占到的appID写入文件
	if _, err = lock.WriteString(appID); err != nil {
		return "", err
	}
	return appID, nil
}

//// Select 为分布式选择器
//// win不支持syscall.Flock，作废，换兼容性更好的
//// 这里使用 "github.com/deptofdefense/safelock" （用起来有点奇怪，作废掉）
//func (s LocalDistributedService) Select(serviceName string, appID string, holdTime time.Duration) (selectAppID string, err error) {
//	appService := s.container.MustMake(contract.AppKey).(contract.App)
//	runtimeFolder := appService.RuntimeFolder()
//
//	filename := "disribute_" + serviceName
//	lockFile := filepath.Join(runtimeFolder, filename)
//
//	// 换成三方库，跨平台兼容性更好
//
//	// 打开文件锁
//	fs := afero.NewOsFs() // 内存：NewMemMapFs()
//	lock := safelock.NewFileLock(0, lockFile, fs)
//
//	// 尝试独占文件锁
//	lockErr := lock.Lock()
//
//	aFile, err := fs.Open(lockFile)
//	if err != nil {
//		if errors.Is(err, afero.ErrFileNotFound) {
//			aFile, err = fs.Create(lockFile)
//		} else {
//			return
//		}
//	}
//	if err != nil {
//		return
//	}
//	defer aFile.Close()
//
//	// 独占失败，读取当前占有的appid
//	if lockErr != nil {
//		selectAppIDByt, readErr := ioutil.ReadAll(aFile)
//		if readErr != nil {
//			return "", readErr
//		}
//		selectAppIDStr := string(selectAppIDByt)
//		return selectAppIDStr, nil
//	}
//
//	// 在一段时间内，选举有效，其他节点在这段时间不能再进行抢占
//	go func() {
//		defer func() {
//			lock.Unlock()
//			fs.Remove(lockFile)
//		}()
//		// 创建选举结果有效的计时器
//		timer := time.NewTimer(holdTime)
//		// 等待计时器结束
//		<-timer.C
//	}()
//
//	// 这里已经是抢占到了，将抢占到的appID写入文件
//	if _, err = aFile.WriteString(appID); err != nil {
//		log.Println(err)
//		return "", err
//	}
//	return appID, nil
//}
