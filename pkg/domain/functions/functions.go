package functions

// Define funções e os métodos para obte-las

import (
	"errors"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"reflect"
	"strings"

	module "github.com/compermane/ic-go/pkg/domain/module"
)

type Function struct {
	Name			string			// Nome da função
	IsMethod		bool			// Se é um método, isto é, se existe um receiver para a função
	ReceiverType 	string
	Module			*module.Module	// Módulo da função
	ArgTypes 		[]string		// Tipos dos argumentos de entrada
	Args 			[]interface{}
	ReturnTypes		[]string		// Tipos das saídas
}

type FunctionSignature map[reflect.Value]Function

// Inicializa uma função com base em suas propriedades
func InitFunction(nome string, is_method bool, receiver_type string, module *module.Module, argTypes, returnTypes []string) (fn *Function, e error) {
	if nome == "" || module == nil {
		e = errors.New("Impossível inicializar uma função com atributos nulos")

		return nil, e
	}

	fn = &Function{
		Name: nome,
		IsMethod: is_method,
		ReceiverType: receiver_type,
		Module: module,
		ArgTypes: argTypes,
		ReturnTypes: returnTypes,
	}

	return fn, nil
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

// Get the functions signatures from a given module. Note that it does not include methods, eg functions declared for a struct/type.
func GetFunctionSignatures(mod *module.Module) (signatures map[string]string, e error) {
	signatures = make(map[string]string)

	for _, mod_path := range mod.Files {
		file, err := os.Open(mod_path)

		if err != nil {
			e = errors.New(fmt.Sprintf("Erro ao abrir arquivo: %v", err))

			return nil, e
		}
		defer file.Close()

		file_set := token.NewFileSet()

		node, err := parser.ParseFile(file_set, mod_path, file, parser.AllErrors)

		if err != nil {
			e = errors.New(fmt.Sprintf("Erro ao analisar o arquivo %v: %v", mod_path, err))

			return nil, e
		}

		conf := types.Config{Importer: importer.Default()}
		info := &types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Defs: make(map[*ast.Ident]types.Object),
		}

		if err != nil {
			e = errors.New(fmt.Sprintf("Erro ao analisar os tipos do arquivo %v: %v\n", mod_path, err))
		}
		
		_, err = conf.Check(mod_path, file_set, []*ast.File{node}, info)
		if err != nil {
			fmt.Printf("Erro ao verificar tipos: %v\n", err)
			return
		}

		ast.Inspect(node, func(n ast.Node) bool {
			if f, ok := n.(*ast.FuncDecl); ok {
				func_name := f.Name.Name

				if obj, ok := info.Defs[f.Name].(*types.Func); ok {
					signature := obj.Type().String()
					signatures[func_name] = signature
				}
			}
			return true
		})
	}

	return signatures, nil
}

func GetFunctionsFromModule(mod *module.Module) (functions []*Function, e error) {
	for _, mod_path := range mod.Files {
		file, err := os.Open(mod_path)

		if err != nil {
			e = errors.New(fmt.Sprintf("Erro ao abrir arquivo: %v", err))

			return nil, e
		}
		defer file.Close()

		file_set := token.NewFileSet()

		node, err := parser.ParseFile(file_set, mod_path, file, parser.AllErrors)

		if err != nil {
			e = errors.New(fmt.Sprintf("Erro ao analisar o arquivo %v: %v", mod_path, err))

			return nil, e
		}

		conf := types.Config{Importer: importer.Default()}
		info := &types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Defs: make(map[*ast.Ident]types.Object),
		}

		if err != nil {
			e = errors.New(fmt.Sprintf("Erro ao analisar os tipos do arquivo %v: %v\n", mod_path, err))
		}
		
		_, err = conf.Check(mod_path, file_set, []*ast.File{node}, info)
		if err != nil {
			fmt.Println("Erro ao verificar tipos:", err)
			return
		}

		ast.Inspect(node, func(n ast.Node) bool {
			if f, ok := n.(*ast.FuncDecl); ok {
				if !strings.Contains(f.Name.Name, "Test") {
					var params, returns []string

					if f.Type.Params != nil {
						for _, param := range f.Type.Params.List {
							for i := 0; i < len(param.Names); i++ {
								param_type := info.TypeOf(param.Type)
								params = append(params, param_type.String())
							}
						}
					}

					if f.Type.Results != nil {
						for _, ret := range f.Type.Results.List {
							ret_type := info.TypeOf(ret.Type)
							returns = append(returns, ret_type.String())
						}
					}

					if f.Recv != nil {
						receiver_type := info.TypeOf(f.Recv.List[0].Type)
						fn, _ := InitFunction(f.Name.Name, true, receiver_type.String(), mod, params, returns)
						functions = append(functions, fn)
					} else {
						fn, _ := InitFunction(f.Name.Name, false, "", mod, params, returns)
						functions = append(functions, fn)
					}
				}
			}
			return true
		})
	}
	return functions, nil
}

func SetFuncArgs(fn *Function) {
	var args []interface{}
	var value interface{}

	for _, tp := range fn.ArgTypes {
		fmt.Println(tp)
		if tp == "float64" {
			value, _ = Float64Generator(1)
		} else if (tp == "float32") {
			value, _ = Float32Generator(1)
		} else if (tp == "int64") {
			value, _ = Int64Generator(1)
		} else if (tp == "int32") {
			value, _ = Int32Generator(1)
		}
		args = append(args, value)
	}

	for _, arg := range args {
		fmt.Println(arg)
	}
	fn.Args = args
}