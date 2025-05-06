package godoparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type FunctionInfo struct {
	FunctionName string
	PackageName  string
	ReceiverName string
	IsStar       bool
	IsMethod     bool
}

func NewFunctionInfo(fn_name, pkg_name, rcv_name string, is_star, is_method bool) FunctionInfo {
	return FunctionInfo{
		fn_name,
		pkg_name,
		rcv_name,
		is_star,
		is_method,
	}
}

func GetPublicFunctionsFromDir(dir_path string) []FunctionInfo {
	funcs        := make([]FunctionInfo, 0)
	entries, err := os.ReadDir(dir_path)

	if err != nil {
		fmt.Printf("Error reading dir path %v: %v\n", dir_path, err)
		return nil
	}
	for _, entry := range entries {
		full_path := filepath.Join(dir_path, entry.Name())

		if entry.IsDir() {
			funcs = append(funcs, GetPublicFunctionsFromDir(full_path)...)
		} else if !strings.HasSuffix(full_path, "_test.go") {
			if filepath.Ext(entry.Name()) != ".go" {
				continue
			}
			funcs = append(funcs, GetPublicFunctionsFromFile(full_path)...)
		}
	}

	return funcs
}

func GetPublicFunctionsFromFile(file_name string) []FunctionInfo {
	funcs := make([]FunctionInfo, 0)
	fset  := token.NewFileSet()
	
	node, err := parser.ParseFile(fset, file_name, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("Error parsing file %v: %v\n", file_name, err)
		return nil
	}

	ast.Inspect(node, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		pkg_name := node.Name.Name
		var rcv_name, fn_name string
		var is_star, is_method bool 

		if isExported(fn.Name.Name) && !strings.HasPrefix(fn.Name.Name, "Test"){
			fn_name = fn.Name.Name

			if fn.Recv != nil && len(fn.Recv.List) >= 0 {
				is_method = true

				switch t := fn.Recv.List[0].Type.(type)  {
				case *ast.Ident:
					is_star = false
					rcv_name = t.Name
				case *ast.StarExpr:
					if ident, ok := t.X.(*ast.Ident); ok {
						is_star  = true
						rcv_name = ident.Name
					}
				}

				if isExported(rcv_name) {
					funcs = append(funcs, NewFunctionInfo(fn_name, pkg_name, rcv_name, is_star, is_method))
				}
				
			} else {
				rcv_name  = ""
				is_star   = false
				is_method = false
				funcs = append(funcs, NewFunctionInfo(fn_name, pkg_name, rcv_name, is_star, is_method))
			}

			return true
		}	
		return false
	})

	return funcs
}

func DumpFunctions(file_name string, finfos []FunctionInfo) {
	str := ""
	file, err := os.OpenFile(file_name, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644) 

	if err != nil {
		panic("Error writing to file " + file_name)
	}

	defer file.Close()

	for _, finfo := range finfos {
		str += finfo.toExecutableFunc()
		str += ",\n"
	}

	file.WriteString(str)
}

func (finfo FunctionInfo) toExecutableFunc() string {
	if finfo.IsMethod && finfo.IsStar {
		return  "(*" + finfo.PackageName + "." + finfo.ReceiverName + ")." + finfo.FunctionName
	}

	if finfo.IsMethod && !finfo.IsStar { 
		return "(" + finfo.PackageName + "." + finfo.ReceiverName + ")." + finfo.FunctionName
	}

	return finfo.PackageName + "." + finfo.FunctionName

}

func isExported(fn_name string) bool {
	if fn_name == "" {
		return false
	}

	return unicode.IsUpper(rune(fn_name[0]))
}