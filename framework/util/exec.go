package util

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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
		return file + string(os.PathSeparator) // Error: daemon: Non-POSIX OS is not supported
	}
	fmt.Println("获取执行目录失败：err=", err.Error())
	return ""
}

// GetBinaryFileDirectory 获取当前执行程序的二进制文件目录
func GetBinaryFileDirectory() string {
	file, err := os.Executable()
	if err == nil {
		return file + string(os.PathSeparator) // Error: daemon: Non-POSIX OS is not supported
	}
	fmt.Println("获取二进制文件目录失败：err=", err.Error())
	return ""
}

// 检查当前执行程序二进制是否在gopath
func CheckBinaryFileInTheGOPATH() bool {
	var homePath = os.Getenv("HOMEPATH")
	binaryPath := GetBinaryFileDirectory()
	goBinPath := filepath.Join(homePath, "go", "bin")
	return strings.Contains(binaryPath, goBinPath)
}

// GetRootDirectory 获取当前项目根目录（根据 .go-root 文件识别）
func GetRootDirectory() (string, error) {
	executable, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := filepath.Dir(executable)
	for {
		if _, err := os.Stat(filepath.Join(dir, ".go-root")); err == nil {
			return dir, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}
	return "", fmt.Errorf("unable to find project root")
}
