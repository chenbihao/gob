package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"
)

// 生成 gob 的 md 文件
func main() {
	genProviderTargetPath, _ := filepath.Abs("./framework/contract")
	genProviderPath, _ := filepath.Abs("./docs/src/provider")

	genCommandTargetPath, _ := filepath.Abs("./framework/command")
	genCommandPath, _ := filepath.Abs("./docs/src/command")

	fmt.Println(fmt.Sprintf("开始生成 服务提供者文档, 目标：%s 生成：%s", genProviderTargetPath, genProviderPath))
	genProviderDocs(genProviderTargetPath, genProviderPath, []string{})

	fmt.Println(fmt.Sprintf("开始生成 命令说明文档, 目标：%s 生成：%s", genCommandTargetPath, genCommandPath))
	genCommandDocs(genCommandTargetPath, genCommandPath, []string{})

	fmt.Println("生成完毕")
}

// 生成服务提供者文档
func genProviderDocs(targetPath, genPath string, selectFiles []string) {
	// get packages
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, targetPath, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("ParseDir 出错:", err)
		return
	}

	// 开始创建文件夹
	if err = CreateFolderIfNotExists(genPath); err != nil {
		fmt.Println("创建文件 出错:", err)
		return
	}

	for _, pkg := range pkgs {
		for filePath, astFile := range pkg.Files {
			fileName := filepath.Base(filePath)
			fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

			if len(selectFiles) != 0 {
				if !slices.Contains(selectFiles, fileName) {
					continue
				}
			}

			fmt.Println("正在生成:", fileName)
			var key, remark, code string

			// 获取说明
			remark = getProviderRemark(filePath, fset, astFile)
			remark = strings.Trim(remark, "/*\n")
			remark = strings.Trim(remark, "\n*/")

			key = getProviderKey(astFile, key)
			code = getProviderCode(astFile, code, filePath, fset)

			file, err := os.Create(filepath.Join(genPath, fileName+".md"))
			if err != nil {
				fmt.Println("创建文件出错:", err)
				return
			}

			p := make(map[string]string)
			p["key"] = key
			p["remark"] = remark
			p["code"] = code

			tpl := template.Must(template.New("").Parse(providerMD))
			if err := tpl.Execute(file, p); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("创建md成功, 文件夹地址:", genPath)
	}
}

// 生成命令说明文档
func genCommandDocs(targetPath, genPath string, selectFiles []string) {
	// get packages
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, targetPath, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("ParseDir 出错:", err)
		return
	}

	// 开始创建文件夹
	if err = CreateFolderIfNotExists(genPath); err != nil {
		fmt.Println("创建文件 出错:", err)
		return
	}

	for _, pkg := range pkgs {
		for filePath, astFile := range pkg.Files {
			fileName := filepath.Base(filePath)
			fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

			if fileName == "kernel" {
				continue
			}
			if len(selectFiles) != 0 {
				if !slices.Contains(selectFiles, fileName) {
					continue
				}
			}

			fmt.Println("正在生成:", fileName)
			var key, remark string

			// 获取说明
			remark = getCommandRemark(filePath, fset, astFile)
			remark = strings.Trim(remark, "/*\n")
			remark = strings.Trim(remark, "\n*/")

			key = getCommandKey(astFile, key)
			//code = getProviderCode(astFile, code, filePath, fset)

			file, err := os.Create(filepath.Join(genPath, fileName+".md"))
			if err != nil {
				fmt.Println("创建文件出错:", err)
				return
			}

			p := make(map[string]string)
			p["key"] = key
			p["remark"] = remark
			p["code"] = ""

			tpl := template.Must(template.New("").Parse(commandMD))
			if err := tpl.Execute(file, p); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("创建md成功, 文件夹地址:", genPath)
	}
}
