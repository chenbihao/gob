package main

import (
	"bufio"
	"bytes"
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
	targetPath, _ := filepath.Abs("./framework/contract")
	genPath, _ := filepath.Abs("./docs/src/provider")
	//genGobDocs(targetPath, genPath, []string{""})
	genGobDocs(targetPath, genPath, []string{})
}

func genGobDocs(targetPath, genPath string, selectFiles []string) {
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

			for _, object := range astFile.Scope.Objects {
				// 获取 key
				if object.Kind == ast.Con {
					d := object.Decl.(*ast.ValueSpec)
					if len(d.Names) > 0 && len(d.Values) > 0 &&
						strings.HasSuffix(d.Names[0].Name, "Key") {
						key = strings.Trim(d.Values[0].(*ast.BasicLit).Value, `"`)
					}
				}
				// 获取 code
				if object.Kind == ast.Typ {
					d := object.Decl.(*ast.TypeSpec)
					switch d.Type.(type) {
					case *ast.InterfaceType:
						// 把源代码截取原文处理
						code = "```go \n" + getSourceCode(filePath, fset, d) + "\n```"
					default:
						// 处理其他类型
						continue
					}

				}
			}

			genFilePath := filepath.Join(genPath, fileName+".md")
			file, err := os.Create(genFilePath)
			if err != nil {
				fmt.Println("创建文件出错:", err)
				return
			}

			tpl := template.Must(template.New("first").Parse(md))
			p := make(map[string]string)
			p["key"] = key
			p["remark"] = remark
			p["code"] = code
			if err := tpl.Execute(file, p); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("创建md成功, 文件夹地址:", genPath)
	}
}

var md = `---
lang: zh-CN
title: {{.key}}
description:
---
# {{.key}}

{{.remark}}

## 使用方法
{{.code}}
`

func getSourceCode(filePath string, fset *token.FileSet, d *ast.TypeSpec) (s string) {
	startLine := fset.Position(d.Type.Pos()).Line
	endLine := fset.Position(d.Type.End()).Line

	f, _ := os.OpenFile(filePath, os.O_RDONLY, 0666)
	defer f.Close()

	buf := make([]byte, 1024*4)
	f.Read(buf)
	reader := bufio.NewReader(bytes.NewReader(buf))

	sb := strings.Builder{}
	for i := 1; i <= endLine; i++ {
		line, _, _ := reader.ReadLine()
		if i >= startLine {
			sb.Write(line)
			if i != endLine {
				sb.Write([]byte("\n"))
			}
		}
	}
	return sb.String()
}

func getRemark(filePath string, fset *token.FileSet, astFile *ast.File) (s string) {

	objects := astFile.Scope.Objects
	var object *ast.Object
	for k, o := range objects {
		if object == nil && strings.HasSuffix(k, "Key") {
			object = o
		}
	}
	if object == nil {
		return
	}

	startLine := 0
	endLine := fset.Position(object.Pos()).Line - 1

	if len(astFile.Imports) != 0 {
		startLine = fset.Position(astFile.Imports[len(astFile.Imports)-1].End()).Line + 1
	} else {
		startLine = fset.Position(astFile.Package).Line + 1
	}

	f, _ := os.OpenFile(filePath, os.O_RDONLY, 0666)
	defer f.Close()

	buf := make([]byte, 1024*4)
	f.Read(buf)
	reader := bufio.NewReader(bytes.NewReader(buf))
	sb := strings.Builder{}
	for i := 1; i < endLine; i++ {
		line, _, _ := reader.ReadLine()
		if i > startLine {
			sb.Write(line)
			if i != endLine {
				sb.Write([]byte("\n"))
			}
		}
	}
	return sb.String()
}

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
