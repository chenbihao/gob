//go:build !windows

package util

import (
	"fmt"
	"os"
	"syscall"

	"github.com/erikdubbelboer/gspt"
)

// SetProcessTitle 设置进程名
func SetProcessTitle(name string) {
	gspt.SetProcTitle(name)
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
