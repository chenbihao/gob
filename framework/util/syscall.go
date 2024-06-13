//go:build !windows

package util

import (
	"os"
	"syscall"
	// "github.com/joho/godotenv"
	// "github.com/erikdubbelboer/gspt"
)

// SetProcessTitle 设置进程名
func SetProcessTitle(name string) {
	// gspt 构建出错 ， 可能只支持mac并且需要安装 xcode-select
	// godotenv.SetProcTitle(name)
}

// CheckProcessExist Will return true if the process with PID exists.
func CheckProcessExist(pid int) bool {
	// 查询这个pid
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	// 给进程发送signal 0, 如果返回nil，代表进程存在, 否则进程不存在
	if err = process.Signal(syscall.Signal(0)); err != nil {
		return false
	}
	return true
}

// KillProcess kill process by pid
func KillProcess(pid int) error {
	return syscall.Kill(pid, syscall.SIGTERM)
}
