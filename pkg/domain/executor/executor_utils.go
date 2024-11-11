package executor

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/compermane/ic-go/pkg/domain/functions"
)

func (exec *Executor) PrintFunctions() {
	for _, fn := range exec.FunctionsList {
		fmt.Println(fn.Name)
	}
}

func (exec *Executor) PrintNextCandidates() {
	for _, fn := range exec.FunctionsList {
		var builder strings.Builder

		if len(exec.NextCandidates[fn]) != 0 {
			for _, fns := range exec.NextCandidates[fn] {
				builder.WriteString(fns.Name + " ")
			}
		} else {
			builder.WriteString("None")
		}
		fmt.Printf("%v: %v\n", fn.Name, builder.String())
	}
}

func (exec *Executor) GetFunctionReturns(fn *functions.Function) string {
	str := "["

	for _, ret := range exec.FunctionsReturns[fn] {
		if ret.Kind() == reflect.Int {
			str = str + " " + fmt.Sprintf("%v", ret.Int())
		} else if ret.Kind() == reflect.Float32 || ret.Kind() == reflect.Float64 {
			str = str + " " +  fmt.Sprintf("%v", ret.Float())
		} else {
			str = str + " " + ret.String()
		}
	}

	str = str + " ]"

	return str
}