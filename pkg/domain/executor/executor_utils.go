package executor

import (
	"fmt"
	"strings"
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