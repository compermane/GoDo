package executor

import (
	"fmt"
)

func (exec *Executor) PrintFunctions() {
	for _, fn := range exec.FunctionsList {
		fmt.Println(fn.Name)
	}
}
