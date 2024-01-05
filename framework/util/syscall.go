package util

import (
	"fmt"
	"os"
	"syscall"
)

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
