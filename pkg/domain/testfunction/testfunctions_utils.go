package testfunction

import "fmt"

func (fn *TestFunction) Print() {
	str_arg_types := ""
	str_args := ""
	str_ret_types := ""

	for i := range fn.ArgTypes {
		str_arg_types = str_arg_types + fn.ArgTypesString[i] + ", "

		if i < len(fn.ArgValues) {
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