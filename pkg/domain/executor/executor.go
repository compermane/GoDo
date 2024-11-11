package executor

import (
	"fmt"
	"reflect"
	"time"

	functions "github.com/compermane/ic-go/pkg/domain/functions"
	"github.com/compermane/ic-go/pkg/domain/receiver"
)

type Executor struct {
	ReceiversList		[]*receiver.Receiver
	FunctionsList		[]*functions.Function
	FunctionsReturns 	map[*functions.Function][]reflect.Value
	NextCandidates		map[*functions.Function][]*functions.Function
	NeededFunctions		map[*functions.Function][]*functions.Function
	FunctionErrors		map[*functions.Function][]any
}

/*  Initializes a function executor.
 *  :param fn_lst: List of functions to be executed
 *  :returns: Pointer to a executor
 */ 
func InitExecutor(fn_lst []any, rcvs_lst []any) *Executor {
	lst := make([]*functions.Function, 0)
	rcvs := make([]*receiver.Receiver, 0)

	for _, fn := range fn_lst {
		lst = append(lst, functions.GetFunction(fn))
	}

	for _, rcv := range rcvs_lst {
		rcvs = append(rcvs, receiver.GetReceiver(rcv))
	}

	return &Executor{
		ReceiversList: rcvs,
		FunctionsList: lst,
		FunctionsReturns: make(map[*functions.Function][]reflect.Value),
		NextCandidates: make(map[*functions.Function][]*functions.Function),
		FunctionErrors: make(map[*functions.Function][]any),
	}
}

func (exec *Executor) ExecuteFunc(fn *functions.Function, args []reflect.Value) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Exception during execution of %v: %v\n", fn.Name, r)
			exec.FunctionErrors[fn] = append(exec.FunctionErrors[fn], r)
		}
	}()

	exec.FunctionsReturns[fn] = fn.Signature.Call(args)
}

func (exec *Executor) ExecuteFuncWithGivenValues(fn *functions.Function, params []reflect.Value) ([]reflect.Value) {
	defer func ()  {
		if r := recover(); r != nil {
			fmt.Printf("Erro durante execução de %v: %v\n", fn.Name, r)
		}
	}()
	
    for i, param := range params {
        if !param.IsValid() {
            panic(fmt.Sprintf("Param %d is invalid (zero Value argument)", i))
        }
    }
	returns := fn.Signature.Call(params)

	return returns
}

func (exec *Executor) GetNextCandidates() {
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

func (exec *Executor) SimpleExecution() {
	for _, fn := range exec.FunctionsList {
		if !fn.IsMethod {
			functions.SetFuncArgs(fn)
			reflect_args := functions.ArgToReflectValue(fn.Args)

			exec.ExecuteFunc(fn, reflect_args)

			fmt.Printf("Function %v ( %v ) succesfully executed: %v\n", fn.Name, fn.Args, exec.GetFunctionReturns(fn))
		} else {
			for _, rcv := range exec.ReceiversList {
				fmt.Println(rcv)
				// rcv_type := reflect.TypeOf(rcv)

				// if rcv_t {

				// }
			}
		}
	}
}

func ExecuteFuncs(fns, rcvs []any, algorithm string, no_runs, timeout int) {
	exec := InitExecutor(fns, rcvs)

	switch algorithm {
	case "baseline1":
		if no_runs != 0 {
			for i := 0; i < no_runs; i++ {
				exec.SimpleExecution()
			}
		} else if timeout != 0 {
			timeout := time.Duration(timeout) * time.Second
			timer := time.After(timeout)
			
			n := 0
			for {
				select {
				case <-timer:
					fmt.Printf("Executed %v funcs\n\n", n)
					return
				default:
					exec.SimpleExecution()
					n = n + 1
				}
			}
		} else {
			panic("Invalid number of runs or timeout duration\n")
		}
	default:
		panic("No algorithm named: " + algorithm + "\n")
	}
}