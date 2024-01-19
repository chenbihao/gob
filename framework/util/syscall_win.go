//go:build windows

package util

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// SetProcessTitle 设置进程名
func SetProcessTitle(name string) {
	// win下无法设置进程名
	return
}

// KillProcess
func KillProcess(pid int, signal syscall.Signal) (err error) {
	// 获取进程信息
	if _, err = os.FindProcess(pid); err != nil {
		fmt.Printf("Failed to find process: %v\n", err)
		return
	}
	return WinKillProcess(pid)
}

func WinKillProcess(pid int) error {
	cmd := exec.Command("taskkill", "/F", "/PID", fmt.Sprintf("%d", pid))
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}
	return nil
}
