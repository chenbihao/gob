package main

import (
	"bufio"
	"bytes"
	"go/ast"
	"go/token"
	"os"
	"strings"
)

func getProviderKey(astFile *ast.File, key string) string {
	for _, object := range astFile.Scope.Objects {
		// 获取 key
		if object.Kind == ast.Con {
			d := object.Decl.(*ast.ValueSpec)
			if len(d.Names) > 0 && len(d.Values) > 0 &&
				strings.HasSuffix(d.Names[0].Name, "Key") {
				key = strings.Trim(d.Values[0].(*ast.BasicLit).Value, `"`)
			}
		}
	}
	return key
}

func getProviderCode(astFile *ast.File, code string, filePath string, fset *token.FileSet) string {
	for _, object := range astFile.Scope.Objects {
		// 获取 code
		if object.Kind == ast.Typ {
			d := object.Decl.(*ast.TypeSpec)
			switch d.Type.(type) {
			case *ast.InterfaceType:
				// 把源代码截取原文处理
				code = "```go \n" + getProviderSourceCode(filePath, fset, d) + "\n```"
			default:
				// 处理其他类型
				continue
			}
		}
	}
	return code
}

func getProviderSourceCode(filePath string, fset *token.FileSet, d *ast.TypeSpec) (s string) {
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

func getProviderRemark(filePath string, fset *token.FileSet, astFile *ast.File) (s string) {
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

func getCommandKey(astFile *ast.File, key string) string {
	for _, object := range astFile.Scope.Objects {
		// 获取 key
		if object.Kind == ast.Con {
			d := object.Decl.(*ast.ValueSpec)
			if len(d.Names) > 0 && len(d.Values) > 0 &&
				strings.HasSuffix(d.Names[0].Name, "Key") {
				key = strings.Trim(d.Values[0].(*ast.BasicLit).Value, `"`)
			}
		}
	}
	return key
}

func getCommandRemark(filePath string, fset *token.FileSet, astFile *ast.File) (s string) {
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
