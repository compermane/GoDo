package godoparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type FunctionInfo struct {
	FilePath     string
	FunctionName string
	PackageName  string
	ReceiverName string
	IsStar       bool
	IsMethod     bool
}

func NewFunctionInfo(file_path, fn_name, pkg_name, rcv_name string, is_star, is_method bool) FunctionInfo {
	return FunctionInfo{
		file_path,
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
					funcs = append(funcs, NewFunctionInfo(file_name, fn_name, pkg_name, rcv_name, is_star, is_method))
				}
				
			} else {
				rcv_name  = ""
				is_star   = false
				is_method = false
				funcs = append(funcs, NewFunctionInfo(file_name, fn_name, pkg_name, rcv_name, is_star, is_method))
			}

			return true
		}	
		return false
	})

	return funcs
}

func GetModsFromRepo(repo_path, dump_file string) {
	mods := make(map[string]bool, 0)

	err := filepath.WalkDir(repo_path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			return nil
		}

		entries, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		for _, e := range entries {
			if !e.IsDir() && strings.HasSuffix(e.Name(), ".go") {
				rel, err := filepath.Rel(repo_path, path)
				if err != nil {
					return err
				}

				mods[rel] = true
				break
			}
		}

		return nil
	})

	file, err := os.OpenFile(dump_file, os.O_TRUNC | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		panic("Error writing to file " + dump_file)
	}

	defer file.Close()

	str := ""

	for mod := range mods {
		str += mod + "\n"
	}

	file.WriteString(str)

}

func GetStructsFromRepo(repo_path, dump_file string) {
	structs := make([]string, 0)

	err := filepath.Walk(repo_path, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			fmt.Println("Erro ao parsear:", path, err)
			return err
		}
	
		pkgName := node.Name.Name

		for _, decl := range node.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok != token.TYPE {
				continue
			}
			for _, spec := range genDecl.Specs {
				typeSpec := spec.(*ast.TypeSpec)
				if !isExported(typeSpec.Name.Name) {
					continue // Ignora structs privadas
				}
				if _, ok := typeSpec.Type.(*ast.StructType); ok {
					structs = append(structs, fmt.Sprintf("%s.%s{},\n", pkgName, typeSpec.Name.Name))
					structs = append(structs, fmt.Sprintf("&%s.%s{},\n", pkgName, typeSpec.Name.Name))
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	DumpToFile(dump_file, structs)
}

func DumpToFile(file_name string, strs []string)  {
	file, err := os.OpenFile(file_name, os.O_TRUNC | os.O_CREATE | os.O_WRONLY, 0644) 

	if err != nil {
		panic("Error writing to file " + file_name)
	}

	defer file.Close()
	
	for _, str := range strs {
		file.WriteString(str)
	}
}
func DumpFunctions(file_name string, finfos []FunctionInfo) {
	str := ""
	file, err := os.OpenFile(file_name, os.O_TRUNC | os.O_CREATE | os.O_WRONLY, 0644) 

	if err != nil {
		panic("Error writing to file " + file_name)
	}

	defer file.Close()

	for _, finfo := range finfos {
		str += finfo.toExecutableFunc()
		str += "\n"
	}

	file.WriteString(str)
}

func (finfo FunctionInfo) toExecutableFunc() string {
	if finfo.IsMethod && finfo.IsStar {
		return  "(*" + finfo.PackageName + "." + finfo.ReceiverName + ")." + finfo.FunctionName +  ",\t\t\t\t //" + finfo.FilePath
	}

	if finfo.IsMethod && !finfo.IsStar { 
		return "(" + finfo.PackageName + "." + finfo.ReceiverName + ")." + finfo.FunctionName + ",\t\t\t\t //" + finfo.FilePath
	}

	return finfo.PackageName + "." + finfo.FunctionName + ",\t\t\t\t //" + finfo.FilePath

}

func isExported(fn_name string) bool {
	if fn_name == "" {
		return false
	}

	return unicode.IsUpper(rune(fn_name[0]))
}