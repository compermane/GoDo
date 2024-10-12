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
	"runtime"
	"strings"

	module "github.com/compermane/ic-go/pkg/domain/module"
	receiver "github.com/compermane/ic-go/pkg/domain/receiver"
	utils "github.com/compermane/ic-go/pkg/utils"
)

type Function struct {
	Name			string				// Nome da função
	Signature 		reflect.Value		// Função propriamente dita a ser executada
	IsMethod		bool				// Se é um método, isto é, se existe um receiver para a função
	Receiver		*receiver.Receiver
	Module			*module.Module		// Módulo da função
	ArgTypes 		[]string			// Tipos dos argumentos de entrada
	Args 			[]interface{}
	ReturnTypes		[]string			// Tipos das saídas
}

// Inicializa uma função com base em suas propriedades
func InitFunction(nome string, is_method bool, receiver_type *receiver.Receiver, module *module.Module, argTypes, returnTypes []string) (fn *Function) {
	fn = &Function{
		Name: nome,
		IsMethod: is_method,
		Receiver: receiver_type,
		Module: module,
		ArgTypes: argTypes,
		ReturnTypes: returnTypes,
	}

	return fn
}

func GetFunction(fn any) *Function {
	fn_value := reflect.ValueOf(fn)
	fn_type := fn_value.Type()
	fn_ptr := fn_value.Pointer()
	fn_info := runtime.FuncForPC(fn_ptr)

	var arg_types, return_types []string
	is_method := false
	name := fn_info.Name()
	for i := 0; i < fn_type.NumIn(); i++ {
		if (i == 0 && fn_type.In(0).Kind() == reflect.Struct) {
			is_method = true
		} else {
			arg_types = append(arg_types, fn_type.In(i).String())
		}
	}

	for i := 0; i < fn_type.NumOut(); i++ {
		return_types = append(return_types, fn_type.Out(i).String())
	}

	function := InitFunction(name, is_method, nil, nil, arg_types, return_types)
	function.Signature = fn_value

	return function
}


func GetFunctionsInformation(mod *module.Module, fns_map map[string]any) (functions []*Function, e error) {
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
						var rcvr *receiver.Receiver

						rcvs, _ := receiver.GetReceivers(mod)
						receiver_name := ""

						for _, field := range f.Recv.List {
							switch expr := field.Type.(type) {
								case *ast.StarExpr: // Ponteiro
									if ident, ok := expr.X.(*ast.Ident); ok {
										receiver_name = ident.Name
									}
								case *ast.Ident: // Não é ponteiro
									receiver_name = expr.Name
							}
						}

						for _, rcv := range rcvs {
							if rcv.Name == receiver_name {
								rcvr = rcv
								break
							}
						}

						receiver.SetAttrValues(rcvr)

						fn := InitFunction(f.Name.Name, true, rcvr, mod, params, returns)
						functions = append(functions, fn)
					} else {
						fn := InitFunction(f.Name.Name, false, nil, mod, params, returns)
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
		if tp == "float64" {
			value, _ = utils.Float64Generator()
		} else if (tp == "float32") {
			value, _ = utils.Float32Generator()
		} else if (tp == "int64") {
			value, _ = utils.Int64Generator()
		} else if (tp == "int32") {
			value, _ = utils.Int32Generator()
		}
		args = append(args, value)
	}

	fn.Args = args
}

func ArgToReflectValue(args []interface{}) (r_values []reflect.Value) {
	for _, arg := range args {
		r_values = append(r_values, reflect.ValueOf(arg))
	}

	return r_values
}