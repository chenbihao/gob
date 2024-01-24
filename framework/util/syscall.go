//go:build !windows

package util

import (
	"fmt"
	"os"
	"syscall"
	//"github.com/joho/godotenv"
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

// KillProcess
func KillProcess(pid int, signal syscall.Signal) (err error) {
	// 获取进程信息
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("Failed to find process: %v\n", err)
		return
	}
	// 向进程发送 SIGTERM 信号
	if err = process.Signal(signal); err != nil {
		fmt.Printf("Failed to send signal: %v\n", err)
		return
	}
	return
}
