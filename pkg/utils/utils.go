package utils

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strings"
)

type Function interface{}

func ApplyOperation(f Function, args ...interface{}) []interface{} {
    fn := reflect.ValueOf(f)
    if len(args) != fn.Type().NumIn() {
        panic("Número incorreto de argumentos fornecidos")
    }

    in := make([]reflect.Value, len(args))
    for i, arg := range args {
        in[i] = reflect.ValueOf(arg)
    }

    results := fn.Call(in)
    out := make([]interface{}, len(results))
    for i, result := range results {
        out[i] = result.Interface()
    }

    return out
}

func GetFunctionInformation(f Function) ([]reflect.Type, error) {
	func_type := reflect.TypeOf(f)

	if func_type.Kind() != reflect.Func {
		return nil, errors.New("O argumento f precisa ser uma função")
	}

	var arg_types []reflect.Type

	for i := 0; i < func_type.NumIn(); i++ {
		arg_type := func_type.In(i)
		arg_types = append(arg_types, arg_type)
	}

	return arg_types, nil
}

func GetFunctionReturnType(f Function) (returnTypes []reflect.Type, e error) {
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
				functions = append(functions, f)
			}
		}
		return true
	})

	return functions, nil
}