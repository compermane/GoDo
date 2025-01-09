// Functions for parsing a go module
// TODO: parece muito verboso as funções, será que não dá pra refatorar?

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

// Helps representing a struct method
type StructMethod struct {
	StructName 	string
	MethodName 	string
}

type ModuleInfo struct {
	Functions	[]string
	Structs		[]string
	Methods		[]StructMethod
	PackageName	string
}

/* Get all public functions from a file
 * param file_path: relative path to the file
 * returns: functions name in a string array
 */
func ExtractFunctionsFromFile(file_path string) []string {
	file_info, err := os.Stat(file_path)

	if err != nil {
		fmt.Errorf("Erro ao acessar arquivo %v: %v\n", file_path, err)
		return nil
	}

	if file_info.IsDir() {
		fmt.Errorf("Caminho %v é um direório\n")
		return nil
	}

	f_set := token.NewFileSet()

	node, err := parser.ParseFile(f_set, file_path, nil, parser.AllErrors)
	if err != nil {
		fmt.Errorf("Erro ao processar arquivo: %v\n", err)
		return nil
	}

	public_funcs := make([]string, 0)

	ast.Inspect(node, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			// Ignore methods
			if fn.Recv != nil {
				return true
			}

			if len(fn.Name.Name) > 0 && unicode.IsUpper(rune(fn.Name.Name[0])) {
				public_funcs = append(public_funcs, fn.Name.Name)
			}
		}
		return true
	})

	return public_funcs
}

func ExtractFunctionsFromDir(dir_path string) map[string][]string {
	funcs := make(map[string][]string)

	filepath.Walk(dir_path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if filepath.Ext(path) == ".go" {
			public_funcs := ExtractFunctionsFromFile(path)
			if err != nil {
				return err
			}
			if len(public_funcs) > 0 {
				funcs[path] = public_funcs
			}
		}
		return nil
	})

	return funcs
}

/* Retrieves all public structs from a file
 * param file_path: Relative path to the file
 * returns: String array containing the public structs name of that file
 */
func ExtractStructsFromFile(file_path string) []string {
	file_info, err := os.Stat(file_path)

	if err != nil {
		fmt.Errorf("Erro ao acessar arquivo %v: %v\n", file_path, err)
		return nil
	}

	if file_info.IsDir() {
		fmt.Errorf("Caminho %v é um direório\n")
		return nil
	}

	f_set := token.NewFileSet()

	node, err := parser.ParseFile(f_set, file_path, nil, parser.AllErrors)
	if err != nil {
		fmt.Errorf("Erro ao processar arquivo: %v\n", err)
		return nil
	}

	public_structs := make([]string, 0)

	for _, decl := range node.Decls {
		gen_decl, ok := decl.(*ast.GenDecl)

		if !ok || gen_decl.Tok != token.TYPE {
			continue
		}

		for _, spec := range gen_decl.Specs {
			type_spec, ok := spec.(*ast.TypeSpec)

			if !ok {
				continue 
			}

			if _, ok := type_spec.Type.(*ast.StructType); ok {
				if len(type_spec.Name.Name) > 0 && unicode.IsUpper(rune(type_spec.Name.Name[0])) {
					public_structs = append(public_structs, type_spec.Name.Name)
				}
			}
		}
	}
	return public_structs 
}

func ExtractStructsFromDir(dir_path string) map[string][]string {
	structs := make(map[string][]string)

	filepath.Walk(dir_path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if filepath.Ext(path) == ".go" {
			public_funcs := ExtractStructsFromFile(path)
			if err != nil {
				return err
			}
			if len(public_funcs) > 0 {
				structs[path] = public_funcs
			}
		}
		return nil
	})

	return structs
}

func ExtractMethodsFromFile(file_path string) []StructMethod {
	file_info, err := os.Stat(file_path)

	if err != nil {
		fmt.Errorf("Erro ao acessar arquivo %v: %v\n", file_path, err)
		return nil
	}

	if file_info.IsDir() {
		fmt.Errorf("Caminho %v é um direório\n")
		return nil
	}

	f_set := token.NewFileSet()

	node, err := parser.ParseFile(f_set, file_path, nil, parser.AllErrors)
	if err != nil {
		fmt.Errorf("Erro ao processar arquivo: %v\n", err)
		return nil
	}

	methods := make([]StructMethod, 0)

	ast.Inspect(node, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if ok && fn.Recv != nil {
			if len(fn.Name.Name) > 0 && unicode.IsUpper(rune(fn.Name.Name[0])) {
				for _, field := range fn.Recv.List {
					ident, ok := field.Type.(*ast.Ident)

					if !ok {
						star_expr, ok := field.Type.(*ast.StarExpr)
						if ok {
							if ident, ok = star_expr.X.(*ast.Ident); ok {
								methods = append(methods, StructMethod{
									StructName: ident.Name,
									MethodName: fn.Name.Name,
								})
							}
						}
						continue
					}
				}
			}
		} else {
			return true
		}
		return true
	})

	return methods
}

func ExtractMethodsFromDir(dir_path string) map[string][]StructMethod {
	methods := make(map[string][]StructMethod)

	filepath.Walk(dir_path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if filepath.Ext(path) == ".go" {
			public_methods := ExtractMethodsFromFile(path)
			if err != nil {
				return err
			}
			if len(public_methods) > 0 {
				methods[path] = public_methods
			}
		}
		return nil
	})

	return methods
}

func ParseModule(mod_path string) map[string]ModuleInfo {
	result := make(map[string]ModuleInfo)

	filepath.WalkDir(mod_path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(path) == ".go" && !strings.Contains(d.Name(), "test") {
			functions := ExtractFunctionsFromFile(path)
			structs := ExtractStructsFromFile(path)
			methods := ExtractMethodsFromFile(path)
			
			// Get package name
			f_set := token.NewFileSet()
			node, _ := parser.ParseFile(f_set, path, nil, parser.PackageClauseOnly)

			result[path] = ModuleInfo{functions, structs, methods, node.Name.Name}
		}

		return nil
	})

	return result
}

func WriteFile(module_info map[string]ModuleInfo, file_name string) {
	file, err := os.Create(file_name)
	if err != nil {
		fmt.Errorf("Erro ao criar arquivo: %v\n", err)
		return 
	}

	defer file.Close()

	for file_path, value := range module_info {
		_, err := file.WriteString(file_path + ":\n")
		if err != nil {
			fmt.Errorf("Erro ao escrever no arquivo: %v\n", err)
			return 
		}

		file.WriteString("\tFunctions: \n\t")
		funcs := ""
		for i := 0; i < len(value.Functions); i++{
			funcs += value.Functions[i] + "\n\t"
		}
		file.WriteString(funcs + "\n")

		file.WriteString("Structs: \n\t")
		structs := ""
		for i := 0; i < len(value.Structs); i++{
			structs += value.Structs[i] + "\n\t"
		}
		file.WriteString(structs + "\n")

		file.WriteString("Methods: \n\t")
		methods := ""
		for i := 0; i < len(value.Methods); i++{
			methods += value.Methods[i].StructName + "." + value.Methods[i].MethodName + "\n\t"
		}
		file.WriteString(methods + "\n")

		file.WriteString("\n----------------------------------------\n")
	}
}

func JoinPackageAndFunction(module_info ModuleInfo) string {
	str := ""

	for i, func_name := range module_info.Functions {
		str += module_info.PackageName + "." + func_name + ", "
		
		if i % 4 == 0 && i != 0 {
			str += "\n"
		}
	}

	return str
}

func WriteUsableSlice(module_info map[string]ModuleInfo, file_name string) {
	str := ""

	for _, value := range module_info {
		str += JoinPackageAndFunction(value)
	}

	file, err := os.Create(file_name)
	if err != nil {
		fmt.Errorf("Erro ac criar arquivo: %v\n", err)
		return 
	}

	defer file.Close()

	file.WriteString(str)
}