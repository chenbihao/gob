package main

import (
	"fmt"
	"go/ast"
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

	genCommandSub := map[string]string{"model": "model"}

	fmt.Println(fmt.Sprintf("开始生成 服务提供者文档, 目标：%s 生成：%s", genProviderTargetPath, genProviderPath))
	genProviderDocs(genProviderTargetPath, genProviderPath, []string{})

	fmt.Println(fmt.Sprintf("开始生成 命令说明文档, 目标：%s 生成：%s", genCommandTargetPath, genCommandPath))
	genCommandDocs(genCommandTargetPath, genCommandPath, genCommandSub, []string{})

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
			remark = getRemark(filePath, fset, astFile)
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
func genCommandDocs(targetPath, genPath string, genCommandSub map[string]string, selectFiles []string) {
	// get packages
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, targetPath, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("ParseDir 出错:", err)
		return
	}

	// fileName - astFile
	allAstFile := map[string]*ast.File{}
	for _, pkg := range pkgs {
		for filePath, astFile := range pkg.Files {
			allAstFile[filePath] = astFile
		}
	}

	for path, key := range genCommandSub {
		subPkgs, err := parser.ParseDir(fset, filepath.Join(targetPath, path), nil, parser.AllErrors)
		if err != nil {
			fmt.Println("ParseSubDir 出错:", err)
			return
		}
		for _, pkg := range subPkgs {
			for filePath, astFile := range pkg.Files {
				fileName := filepath.Base(filePath)
				fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
				if fileName == key {
					allAstFile[filePath] = astFile
				}
			}
		}
	}

	// 开始创建文件夹
	if err = CreateFolderIfNotExists(genPath); err != nil {
		fmt.Println("创建文件 出错:", err)
		return
	}

	for filePath, astFile := range allAstFile {
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
		remark = getRemark(filePath, fset, astFile)
		remark = strings.Trim(remark, "/*\n")
		remark = strings.Trim(remark, "\n*/")

		key = getCommandKey(astFile, key)

		// todo 执行命令后获取 help ，这里考虑换掉命令行，所以先等命令行换了再来调整
		//code := getCommandDocs(astFile, code, filePath, fset)
		code := "稍后补全，可以使用`gob [command] help`命令获取相关帮助"

		file, err := os.Create(filepath.Join(genPath, fileName+".md"))
		if err != nil {
			fmt.Println("创建文件出错:", err)
			return
		}

		p := make(map[string]string)
		p["key"] = key
		p["remark"] = remark
		p["code"] = code

		tpl := template.Must(template.New("").Parse(commandMD))
		if err := tpl.Execute(file, p); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("创建md成功, 文件夹地址:", genPath)

}
