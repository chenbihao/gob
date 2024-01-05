package util

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
)

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// GetExecDirectory 获取当前执行程序目录
func GetExecDirectory() string {
	file, err := os.Getwd()
	if err == nil {
		//return file + "/"
		return file + string(os.PathSeparator) // Error: daemon: Non-POSIX OS is not supported
	}
	return ""
}

// CheckProcessExist Will return true if the process with PID exists.
func CheckProcessExist(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	err = process.Signal(syscall.Signal(0))
	if err != nil {
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
