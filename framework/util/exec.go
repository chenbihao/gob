package util

import (
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
