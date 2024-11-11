package functions

// Define funções e os métodos para obte-las

import (
	"reflect"
	"runtime"
	"strings"

	utils "github.com/compermane/ic-go/pkg/utils"
)

type Function struct {
	Name			string				// Nome da função
	Signature 		reflect.Value		// Função propriamente dita a ser executada
	IsMethod		bool				// Se é um método, isto é, se existe um receiver para a função
	ReceiverName	string		
	ArgTypes 		[]string			// Tipos dos argumentos de entrada
	Args 			[]interface{}
	ReturnTypes		[]string			// Tipos das saídas
}

// Inicializa uma função com base em suas propriedades
func InitFunction(nome string, is_method bool, receiver_name string, argTypes, returnTypes []string) (fn *Function) {
	fn = &Function{
		Name: nome,
		IsMethod: is_method,
		ReceiverName: receiver_name,
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
	full_name := fn_info.Name()
	
	parts := strings.Split(full_name, ".")
	simple_name := parts[len(parts) - 1]
	name := strings.TrimSuffix(simple_name, "-fm")

	receiver_name := ""
	for i := 0; i < fn_type.NumIn(); i++ {
		if i == 0 && (fn_type.In(0).Kind() == reflect.Struct || fn_type.In(0).Kind() == reflect.Ptr) {
			is_method = true
            receiverType := fn_type.In(0)

            if receiverType.Kind() == reflect.Ptr {
                receiverType = receiverType.Elem()
            }
            receiver_name = receiverType.Name() 
		} else {
			if fn_type.In(i).Kind() == reflect.Struct {
				arg_types = append(arg_types, fn_type.In(i).Name())
			} else if fn_type.In(i).Kind() == reflect.Ptr {
				arg_types = append(arg_types, fn_type.In(i).Elem().String())
			} else {
				arg_types = append(arg_types, fn_type.In(i).String())
			}
		}
	}

	for i := 0; i < fn_type.NumOut(); i++ {
		return_types = append(return_types, fn_type.Out(i).String())
	}

	function := InitFunction(name, is_method, receiver_name, arg_types, return_types)
	function.Signature = fn_value

	return function
}

func SetFuncArgs(fn *Function) {
	var args []interface{}
	var value interface{}

	for _, tp := range fn.ArgTypes {
		if tp == "float64" {
			value, _ = utils.Float64Generator()
		} else if (tp == "float32") {
			value, _ = utils.Float32Generator()
		} else if (tp == "int") {
			value, _ = utils.IntGenerator()
		} else if (tp == "int64") {
			value, _ = utils.Int64Generator()
		} else if (tp == "int32") {
			value, _ = utils.Int32Generator()
		} else if (tp == "string") {
			lenght, _ := utils.IntGenerator(1, 10)
			value = utils.StringGenerator(lenght)
		} else if (tp == "bool") {
			decider, _ := utils.IntGenerator(0, 10)
			value = utils.BooleanGenerator(decider)
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