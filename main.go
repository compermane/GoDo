package main

import (
	"fmt"
	"reflect"

	geometry "github.com/compermane/ic-go/geometry"
	functions "github.com/compermane/ic-go/pkg/domain/functions"
	module "github.com/compermane/ic-go/pkg/domain/module"
	executor "github.com/compermane/ic-go/pkg/executor"
)

func func_info() {
    imports := []string{"math", "fmt"}
    mod, _ := module.InitModule("/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry/geometry.go",
                                "/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry", "github.com/compermane/ic-go/geometry", imports)           

    functions, _ := functions.GetFunctionsFromModule(mod)
    functions = functions[0:1]
    for _, fn := range functions {
        fmt.Printf("%v, %v, Receiver: %v ", fn.Name, fn.IsMethod, fn.ReceiverType)

        fmt.Print("Args: [")
        for _, arg := range fn.ArgTypes {
            fmt.Printf("%v ", arg)
        }
        fmt.Print("] ")

        fmt.Print("Returns: [")
        for _, ret := range fn.ReturnTypes {
            fmt.Printf("%v ", ret)
        }
        fmt.Print("] ")

        print("\n")
    }
}

func func_signatures() {
    imports := []string{"math", "fmt"}
    mod, _ := module.InitModule("/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry/geometry.go",
    "/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry", "github.com/compermane/ic-go/geometry", imports)  

    signatures, _ := functions.GetFunctionSignatures(mod)

    for key, value := range signatures {
        fmt.Printf("%v: %v\n", key, value)
    }
}

func func_executor() {
    var args []reflect.Value
    fn := reflect.ValueOf(geometry.InitPoint)

    args = append(args, reflect.ValueOf(10.0))
    args = append(args, reflect.ValueOf(10.0))

    results := fn.Call(args)

    for _, result := range results {
        fmt.Println(result.Type())
        fmt.Println(result)
    }
    imports := []string{"math", "fmt"}
    mod, _ := module.InitModule("/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry/geometry.go",
                                "/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry", "github.com/compermane/ic-go/geometry", imports)      
    fns, _ := functions.GetFunctionSignatures(mod)

    for key, value := range fns {
        fmt.Printf("%v: %v\n", key, value)
    }

    // float64_values, _ := functions.Float64Generator(2, -10, 20)
    // reflect_values := functions.Float64ToReflectValues(float64_values)
    // result := functions.RunFunc(fns["InitPoint"], reflect_values)

    // fmt.Println(result)
}

func f_utils() {
    values, _ := functions.Float64Generator(3, -1000, 10)

    fmt.Println(values)
}

func exec_test() {
    imports := []string{"math", "fmt"}
    mod, _ := module.InitModule("/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry/geometry.go",
    "/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry", "github.com/compermane/ic-go/geometry", imports)  

    fns, _ := functions.GetFunctionsFromModule(mod)

    fns = fns[0:1]
    functions.SetFuncArgs(fns[0])

    executor.WriteIntermediateFile(fns, mod)
}

func main() {
    exec_test()
    // func_executor()
    // func_info()
    // func_signatures()
    // f_utils()
}