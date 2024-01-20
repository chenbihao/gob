package util

import (
	"fmt"
	"os"
	"runtime"
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
