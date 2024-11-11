package functions

import "fmt"

/*
	Name			string				// Nome da função
	Signature 		reflect.Value		// Função propriamente dita a ser executada
	IsMethod		bool				// Se é um método, isto é, se existe um receiver para a função
	Receiver		*receiver.Receiver
	ArgTypes 		[]string			// Tipos dos argumentos de entrada
	Args 			[]interface{}
	ReturnTypes		[]string			// Tipos das saídas
*/

func (fn *Function) Print() {
	str_arg_types := ""
	str_args := ""
	str_ret_types := ""

	for i := range fn.ArgTypes {
		str_arg_types = str_arg_types + fn.ArgTypes[i] + ", "

		if i < len(fn.Args) {
			str_args = str_args + fmt.Sprintf("%v, ", str_args[i])
		}
	}

	for i := range fn.ReturnTypes {
		str_ret_types = str_ret_types + fn.ReturnTypes[i]
	}

	fmt.Println("Func name: " + fn.Name)
	fmt.Println("Is method: " + fmt.Sprintf("%v", fn.IsMethod))
	fmt.Println("Receiver name: " + fmt.Sprintf("%v", fn.ReceiverName))
	fmt.Println("Args types: " + str_arg_types)
	fmt.Println("Args values: " + str_args)
	fmt.Println("Return types: " + str_ret_types)
}