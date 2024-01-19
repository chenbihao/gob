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
func IsNotWindows() bool {
	return !IsWindows()
}

// GetExecDirectory 获取当前执行程序目录
func GetExecDirectory() string {
	file, err := os.Getwd()
	if err == nil {
		//return file + "/"
		return file + string(os.PathSeparator) // Error: daemon: Non-POSIX OS is not supported
	}
	fmt.Println("获取执行目录失败：err=", err.Error())
	return ""
}

// CheckProcessExist Will return true if the process with PID exists.
func CheckProcessExist(pid int) bool {
	// 查询这个pid
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	if IsNotWindows() {
		// 给进程发送signal 0, 如果返回nil，代表进程存在, 否则进程不存在
		if err = process.Signal(syscall.Signal(0)); err != nil {
			return false
		}
	}
	return true
}
