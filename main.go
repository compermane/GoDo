package main

import (
	"fmt"
	"reflect"

	geometry "github.com/compermane/ic-go/geometry"
	executor "github.com/compermane/ic-go/pkg/domain/executor"
	functions "github.com/compermane/ic-go/pkg/domain/functions"
	module "github.com/compermane/ic-go/pkg/domain/module"
	receiver "github.com/compermane/ic-go/pkg/domain/receiver"
)

func func_executor() {
    var fns map[string]any

    fns["InitPoint"] = geometry.InitPoint
    fns["Add"] = geometry.Add
    var args []reflect.Value
    fn := reflect.ValueOf(geometry.InitPoint)

    fmt.Printf("value of: %v\n", reflect.ValueOf(fn))
    args = append(args, reflect.ValueOf(10.0))
    args = append(args, reflect.ValueOf(10.0))

    results := fn.Call(args)

    for _, result := range results {
        fmt.Println(result.Type())
        fmt.Println(result)
    }
}

func info_with_reflect() {
    fn := geometry.InitPoint

    function := functions.GetFunction(fn)

    fmt.Printf("name: %v\n", function.Name)
    for _, arg := range function.ArgTypes {
        fmt.Printf("%v ", arg)
    }
    fmt.Println()

    for _, ret := range function.ReturnTypes {
        fmt.Printf("%v ", ret)
    }

    functions.SetFuncArgs(function)
    r_args := functions.ArgToReflectValue(function.Args)
    results := function.Signature.Call(r_args)

    for _, result := range results {
        fmt.Println(result)
    }
}

func receivers_test() {
    imports := []string{"math", "fmt"}
    mod, _ := module.InitModule("/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry/geometry.go",
    "/home/eugenio/Área de trabalho/Grad/IC/ic-go/geometry", "github.com/compermane/ic-go/geometry", imports)  

    structs, _ := receiver.GetReceivers(mod)

    for _, struct_rcv := range structs {
        fmt.Printf("Name: %v\n", struct_rcv.Name)
        for i := 0; i < len(struct_rcv.AttrNames); i++ {
            fmt.Printf("%v: %v\n", struct_rcv.AttrNames[i], struct_rcv.AttrTypes[i])
        }
    }
}

func executor_test() {
    funcs := []any{geometry.InitPoint, geometry.GetLineFromPoints, geometry.Add}

    executor := executor.InitExecutor(funcs)

    executor.AnalyseFuncs()

    executor.PrintFunctions()
    executor.PrintNextCandidates()
}

func main() {
    // exec_test()
    // receivers_test()
    // func_executor()
    // info_with_reflect(
    executor_test()
    // func_info()
    // func_signatures()
    // f_utils()
}