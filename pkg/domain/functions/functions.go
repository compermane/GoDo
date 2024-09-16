package funcions

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strings"

	module "github.com/compermane/ic-go/pkg/domain/module"
)

type Function struct {
	Name		string
	Module		*module.Module
	ArgTypes 	[]reflect.Type
	ReturnTypes	[]reflect.Type
}

func InitFunction(nome string, module *module.Module, argTypes, returnTypes []reflect.Type) (fn *Function, e error) {
	if nome == "" || module == nil || argTypes == nil || returnTypes == nil {
		e = errors.New("Impossível inicializar uma função com atributos nulos")

		return nil, e
	}

	fn = &Function{
		Name: nome,
		Module: module,
		ArgTypes: argTypes,
		ReturnTypes: returnTypes,
	}

	return fn, nil
}

func GetFunctionNames(file_path string) (names []string, e error) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, file_path, nil, 0) 

	if err != nil {
		e = errors.New(fmt.Sprintf("Erro ao processar arquivo %v: %v\n", file_path, err))
		return nil, e
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if func_decl, ok := n.(*ast.FuncDecl); ok {
			names = append(names, func_decl.Name.Name)
		}

		return true
	})

	return names, nil
}

func GetArgsTypes(f interface{}) ([]reflect.Type, error) {
	func_type := reflect.TypeOf(f)

	if func_type.Kind() != reflect.Func {
		return nil, errors.New(fmt.Sprintf("O argumento f precisa ser uma função. Recebeu: %v", func_type.Kind()))
	}

	var arg_types []reflect.Type

	for i := 0; i < func_type.NumIn(); i++ {
		arg_type := func_type.In(i)
		arg_types = append(arg_types, arg_type)
	}

	return arg_types, nil
}

func GetFunctionReturnType(f interface{}) (returnTypes []reflect.Type, e error) {
	funcType := reflect.TypeOf(f)

	if funcType.Kind() != reflect.Func {
		e = errors.New("Argumento f deveria ser uma função")
		return nil, e
	}

	for i := 0; i < funcType.NumOut(); i++ {
		returnTypes = append(returnTypes, funcType.Out(i))
	}

	return returnTypes, nil;
}

func GetFunctionsFromFile(file_path string) (functions []Function, e error) {
	file, err := os.Open(file_path)

	if err != nil {
		e = errors.New(fmt.Sprintf("Erro ao abrir arquivo: %v", err))

		return nil, e
	}
	defer file.Close()

	file_set := token.NewFileSet()

	node, err := parser.ParseFile(file_set, file_path, file, parser.AllErrors)

	if err != nil {
		e = errors.New(fmt.Sprintf("Erro ao analisar o arquivo %v: %v", file_path, err))

		return nil, e
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if f, ok := n.(*ast.FuncDecl); ok {
			if !strings.Contains(f.Name.Name, "Test") {
				functions = append(functions, Function{f.Name.Name, node.Name.Name})
			}
		}
		return true
	})

	return functions, nil
}