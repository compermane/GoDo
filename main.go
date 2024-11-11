package main

import (
	geometry "github.com/compermane/ic-go/geometry"
	executor "github.com/compermane/ic-go/pkg/domain/executor"
	"github.com/compermane/ic-go/pkg/domain/functions"
	"github.com/compermane/ic-go/pkg/domain/receiver"
)

func executor_test() {
    // funcs := []any{geometry.Add, geometry.MinusOne, geometry.InitPoint, geometry.GetLineFromPoints, geometry.PointToLineDistance, 
                //    geometry.AlwaysPanic, geometry.SumString}
    funcs := []any{geometry.BoolFunc}
    rcvs := []any{geometry.Point{}}

    executor.ExecuteFuncs(funcs, rcvs, "baseline1", 67250, 10)

    // executor.PrintFunctions()
    // executor.PrintNextCandidates()
}

func receivers_test() {
    rcv := receiver.GetReceiver(&geometry.All{})
    rcv.SetReceiverValues()

    rcv.Print()
}

func functions_test() {
    // point := geometry.Point{}
    // fn := {}any[geometry.Add, geo]
    // fn := functions.GetFunction([geom)
    fn := functions.GetFunction(geometry.InitPoint)

    fn.Print()
}

func main() {
    // executor_test()
    receivers_test()
    // functions_test()
}