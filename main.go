package main

import (
	"bytes"

	geometry "github.com/compermane/ic-go/geometry"
	executor "github.com/compermane/ic-go/pkg/domain/executor"
	"github.com/compermane/ic-go/pkg/domain/functions"
	"github.com/compermane/ic-go/pkg/domain/receiver"
	godoparser "github.com/compermane/ic-go/pkg/godo_parser"
)

func executor_test() {
    // funcs := []any{geometry.Add, geometry.MinusOne, geometry.InitPoint, geometry.GetLineFromPoints, geometry.PointToLineDistance, 
                //    geometry.AlwaysPanic, geometry.SumString}
    funcs := []any{geometry.Example}
    rcvs := []any{geometry.Point{/*X: 1.0, Y: 2.0 */}}

    exec := executor.InitExecutor(funcs, rcvs)

    // fn := functions.GetFunction(geometry.Example)
    // exec.SetFuncArgs(fn)
    exec.SimpleExecution()
    // executor.ExecuteFuncs(funcs, rcvs, "baseline1", 67250, 10)

    // executor.PrintFunctions()
    // executor.PrintNextCandidates()
}

func receivers_test() {
    // rcv := receiver.GetReceiver(geometry.All{})
    // rcv.SetReceiverValues()
    rcv := receiver.GetReceiver(bytes.Buffer{})
    rcv.SetReceiverValues()
    // exec := executor.InitExecutor([]any{}, []any{rcv})

    // exec.ExecuteMethod(rcv, "ExampleWithArgs")
    rcv.Print()
}

func functions_test() {
    // point := geometry.Point{}
    // fn := {}any[geometry.Add, geo]
    // fn := functions.GetFunction([geom)
    fn := functions.GetFunction(geometry.InitPoint)

    fn.Print()
}

func parser_test() {
    // fmt.Println("FUNCTIONS")
    // p_funcs := godoparser.ExtractFunctionsFromFile("./pkg/domain/functions/functions.go")
    // dir_funcs := godoparser.ExtractFunctionsFromDir("./pkg/domain/functions")
    // fmt.Println(p_funcs)
    // fmt.Println(dir_funcs)

    // fmt.Println("STRUCTS")
    // p_structs := godoparser.ExtractStructsFromFile("./pkg/domain/functions/functions.go")
    // dir_structs := godoparser.ExtractStructsFromDir("./pkg/domain/functions")
    // fmt.Println(p_structs)
    // fmt.Println(dir_structs)

    // fmt.Println("METHODS")
    // p_mehtods := godoparser.ExtractMethodsFromFile("./pkg/domain/executor/executor.go")
    // dir_methods := godoparser.ExtractMethodsFromDir("./pkg/domain/executor")
    // fmt.Println(p_mehtods)
    // fmt.Println(dir_methods)

    // fmt.Println("MODULE")
    domain := godoparser.ParseModule("./testrepos/cobra")
    // fmt.Println(domain)
    godoparser.WriteUsableSlice(domain, "bruh")
    // godoparser.WriteFile(domain, "cobra_info.txt")
}

func main() {
    // executor_test()
    // receivers_test() 
    // functions_test()
    parser_test()
}