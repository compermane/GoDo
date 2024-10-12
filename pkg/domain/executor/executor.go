package executor

import (
	"reflect"

	functions "github.com/compermane/ic-go/pkg/domain/functions"
)


type Executor struct {
	FunctionsList		[]*functions.Function
	FunctionsReturns 	map[*functions.Function][]reflect.Value
	NextCandidates		map[*functions.Function][]*functions.Function
}

/*  Initializes a function executor.
 *  :param fn_lst: List of functions to be executed
 *  :returns: Pointer to a executor
 */ 
func InitExecutor(fn_lst []any) *Executor {
	lst := make([]*functions.Function, 0)

	for _, fn := range fn_lst {
		lst = append(lst, functions.GetFunction(fn))
	}

	return &Executor{
		FunctionsList: lst,
		FunctionsReturns: make(map[*functions.Function][]reflect.Value),
		NextCandidates: make(map[*functions.Function][]*functions.Function),
	}
}

/*  Execute a function, giving it random values based on its arguments.
 *  :param fn: Function to be executed
 *  :returns: Slice containing returned values
 */
func ExecuteFunc(fn *functions.Function) ([]reflect.Value) {
	// First, set function params
	functions.SetFuncArgs(fn)

	// Convert arguments types to reflect types
	reflect_args := functions.ArgToReflectValue(fn.Args)

	// Execute with reflect and get its results
	returns := fn.Signature.Call(reflect_args)	

	return returns
}

/*  Execute each declared function from a given executor
 *  :returns: None
 */
func (exec *Executor) ExecuteFuncs() {
	// Execute each function previously declared
	for _, fn := range exec.FunctionsList {
		returns := ExecuteFunc(fn)

		// Append its returned value to the map
		exec.FunctionsReturns[fn] = returns
	}
}

func (exec *Executor) AnalyseFuncs() {
	for _, fn := range exec.FunctionsList {
		for _, fn_next := range exec.FunctionsList {
			for _, ret_type := range fn.ReturnTypes {
				for _, arg_type := range fn_next.ArgTypes {
					if ret_type == arg_type && fn.Signature != fn_next.Signature {
						exec.NextCandidates[fn] = append(exec.NextCandidates[fn], fn_next)
						break
					}
				}
			}
		}
	}
}