//go:build windows

package util

import (
	"fmt"
	"os"
)

// SetProcessTitle 设置进程名
func SetProcessTitle(name string) {
	// win下无法设置进程名
	return
}

// CheckProcessExist Will return true if the process with PID exists.
func CheckProcessExist(pid int) bool {
	// 查询这个pid
	if _, err := os.FindProcess(pid); err != nil {
		return false
	}
	return true
}

// KillProcess
func KillProcess(pid int) (err error) {
	// 获取进程信息
	p, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("Failed to find process: %v\n", err)
		return err
	}
	return p.Kill()
}
