package godo

import (
	"fmt"

	utils "github.com/compermane/ic-go/pkg/utils"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	arg_types, err := utils.GetFunctionInformation(Add)

	if err != nil {
		fmt.Printf("Erro: %v", err)

		return
	}

	fmt.Printf("Tipos: %v", arg_types)
}