package functions

// Define funções e os métodos para obte-las

import (
	"reflect"
	"runtime"
	"strings"
)

type Function struct {
	Name			string				// Nome da função
	Signature 		reflect.Value		// Função propriamente dita a ser executada
	IsMethod		bool				// Se é um método, isto é, se existe um receiver para a função
	HasVariadic		bool				// True se o último argumento for variádico, false caso contrário
	ReceiverName	string				// Nome do receiver, se tiver. Pode ser ""
	ArgTypesString 	[]string			// Tipos dos argumentos de entrada em string
	ArgTypes		[]reflect.Type      // Tipos dos argumentos de entrada em reflect.Type
	ReturnTypes		[]string			// Tipos das saídas
}

type Method struct {
	Name 			string
	ArgTypes		[]string
	ReturnTypes		[]string
}

// Inicializa uma função com base em suas propriedades
func InitFunction(nome string, receiver_name string, has_variadic bool, arg_types []reflect.Type, arg_types_string, returnTypes []string) (fn *Function) {
	fn = &Function{
		Name: nome,
		HasVariadic: has_variadic,
		ReceiverName: receiver_name,
		ArgTypesString: arg_types_string,
		ArgTypes: arg_types,
		ReturnTypes: returnTypes,
	}

	return fn
}

func GetFunction(fn any) *Function {
	fn_value := reflect.ValueOf(fn)
	fn_type := fn_value.Type()
	fn_ptr := fn_value.Pointer()
	fn_info := runtime.FuncForPC(fn_ptr)

	var arg_types_string, return_types []string
	var arg_types []reflect.Type
	var is_variadic bool = false
	full_name := fn_info.Name()
	
	parts := strings.Split(full_name, ".")
	simple_name := parts[len(parts) - 1]
	name := strings.TrimSuffix(simple_name, "-fm")

	receiver_name := ""
	for i := 0; i < fn_type.NumIn(); i++ {
		if i == fn_type.NumIn() - 1 {
			if fn_type.IsVariadic() {
				is_variadic = true
			}
		}
		if fn_type.In(i).Kind() == reflect.Struct {
			arg_types_string = append(arg_types_string, fn_type.In(i).Name())
			arg_types = append(arg_types, fn_type.In(i))
		} else if fn_type.In(i).Kind() == reflect.Ptr {
			arg_types_string = append(arg_types_string, fn_type.In(i).Elem().String())
			arg_types = append(arg_types, fn_type.In(i))
		} else {
			arg_types_string = append(arg_types_string, fn_type.In(i).String())
			arg_types = append(arg_types, fn_type.In(i))
		}
		
	}

	for i := 0; i < fn_type.NumOut(); i++ {
		return_types = append(return_types, fn_type.Out(i).String())
	}

	function := InitFunction(name, receiver_name, is_variadic, arg_types, arg_types_string, return_types)
	function.Signature = fn_value


	return function
}
