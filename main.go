package main

import (
	"fmt"
	"os"
	"plugin"
	"reflect"

	utils "github.com/compermane/ic-go/pkg/utils"
)

func main() {
    filePath := os.Args[1]
    p, err := plugin.Open("plugin_geometry.so")

    if err != nil {
        panic(fmt.Sprintf("Erro durante abertura de plugin %v", err))
    }
    functions, err := utils.GetFunctionsFromFile(filePath)

    if err != nil {
        panic(fmt.Sprintf("Erro durante execucao: %v", err))
    }

    for _, function := range functions {
        func_symbol, err := p.Lookup(function.FuncName)

        if err != nil {
            fmt.Printf("Ignorando método %v (TODO: tratar métodos)\n", function.FuncName)
            continue
        }

        arg_types, err := utils.GetFunctionInformation(func_symbol)
        // fn, ok := func_symbol.(func())
        // fmt.Printf("ok: %v\n", ok)
        arg_values, _ := utils.InputFactory(arg_types)

        if err != nil {
            fmt.Printf("Erro durante extração de tipos de argumentos: %v\n", err)
        }

        // return_types, err := utils.GetFunctionReturnType(func_symbol)

        if err != nil {
            fmt.Printf("Erro durante extração de tipos de retorno: %v\n", err)
        }

        fn := reflect.ValueOf(func_symbol)
        results := fn.Call(arg_values)

        for _, result := range results {
            val := reflect.ValueOf(result)
            typ := reflect.TypeOf(result)

            if val.Kind() == reflect.Struct {
                for i := 0; i < val.NumField(); i++ {
                    field := val.Field(i)
                    field_name := typ.Field(i).Name

                    fmt.Printf("(FIELD_NAME) %v: %v\n", field_name, field)
                }
            }
        }
        // fmt.Printf("results: %v\n", results)
        // fmt.Printf("%v, %v, (%v), (%v), (%v)\n", function.FuncName, function.ModName, arg_types, return_types, arg_values)

    }
}