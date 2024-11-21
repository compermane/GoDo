package executor

import (
	"fmt"
	"reflect"
	"time"

	functions "github.com/compermane/ic-go/pkg/domain/functions"
	"github.com/compermane/ic-go/pkg/domain/receiver"
	"github.com/compermane/ic-go/pkg/utils"
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
 *  :returns: Pointer to an executor
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

func (exec *Executor) ExecuteMethod(rcv *receiver.Receiver, method_name string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Exception during execution of method %v: %v\n", method_name, r)
		}
	}()

	obj := reflect.ValueOf(rcv.Receiver)
	method := obj.MethodByName(method_name)
	
	if !method.IsValid() {
		panic(method_name + " method not found")
	}
	
	method_type := method.Type()
	num_args := method_type.NumIn()

	args := make([]reflect.Value, num_args)

	for i := 0; i < num_args; i++ {
		arg_type := method_type.In(i)
		switch arg_type.Kind() {
		case reflect.Int:
			arg, _ := utils.IntGenerator(-100, 100)
			args[i] = reflect.ValueOf(arg)
		case reflect.String:
			lenght, _ := utils.IntGenerator(1, 100)
			args[i] = reflect.ValueOf(utils.StringGenerator(lenght))
		}
	}

	// Talvez eu precise de mais um atributo em executor para armazenar os métodos. 
	results := method.Call(args)

	for _, result := range results {
		fmt.Println(result)
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
		if no_runs > 0 {
			for i := 0; i < no_runs; i++ {
				exec.SimpleExecution()
			}
		} else if timeout > 0 {
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