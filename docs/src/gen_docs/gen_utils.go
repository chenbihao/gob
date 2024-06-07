package main

import "os"

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	// os.Stat 获取文件信息
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
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
