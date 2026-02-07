package util

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 文件和目录权限常量
const (
	DirPerm  os.FileMode = 0755 // 目录权限：rwxr-xr-x
	FilePerm os.FileMode = 0644 // 文件权限：rw-r--r--
)

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	// os.Stat 获取文件信息
	if _, err := os.Stat(path); err != nil {
		return os.IsExist(err)
	}
	return true
}

// 如果不存在，则创建文件夹
func CreateFolderIfNotExists(folder string) error {
	if !Exists(folder) {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// 如果不存在，则创建文件
func CreateFileIfNotExists(file string) error {
	// 检查文件是否存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		// 创建文件
		file, err := os.Create(file)
		if err != nil {
			return err
		}
		defer file.Close()
	} else if err != nil {
		return err
	}
	return nil
}

// 如果不存在，则创建文件（包含目录）
func CreateFolderFileIfNotExists(folder string, file string) error {
	if err := CreateFolderIfNotExists(folder); err != nil {
		return err
	}
	if err := CreateFileIfNotExists(file); err != nil {
		return err
	}
	return nil
}

// 路径是否是隐藏路径
func IsHiddenDirectory(path string) bool {
	return len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".")
}

// 输出所有子目录，目录名
func SubDir(folder string) ([]string, error) {
	subs, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	ret := []string{}
	for _, sub := range subs {
		if sub.IsDir() {
			ret = append(ret, sub.Name())
		}
	}
	return ret, nil
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Get the data
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("warning: failed to close response body: %v", err)
		}
	}()

	// Check response status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("download failed with status code: %d", resp.StatusCode)
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		if err := out.Close(); err != nil {
			log.Printf("warning: failed to close file: %v", err)
		}
	}()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// CopyFolder 将一个目录复制到另外一个目录中
func CopyFolder(source, destination string) error {
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 使用 filepath.Rel 计算相对路径，避免字符串替换的问题
		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		if relPath == "." {
			// 跳过根目录
			return nil
		}
		if info.IsDir() {
			return os.Mkdir(filepath.Join(destination, relPath), DirPerm)
		} else {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			return os.WriteFile(filepath.Join(destination, relPath), data, FilePerm)
		}
	})
}

// CopyFile 将一个目录复制到另外一个目录中
func CopyFile(source, destination string) error {
	var data, err1 = os.ReadFile(source)
	if err1 != nil {
		return err1
	}
	return os.WriteFile(destination, data, 0777)
}
