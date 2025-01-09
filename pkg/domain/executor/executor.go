package executor

import (
	"bytes"
	"fmt"
	"io"
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
	
	fmt.Printf("ARGS: %v", args)
	exec.FunctionsReturns[fn] = fn.Signature.Call(args)
	fmt.Printf("Function %v ( %v ) succesfully executed: %v\n", fn.Name, fn.Args, exec.GetFunctionReturns(fn))
}

func (exec *Executor) ExecuteMethod(rcv *receiver.Receiver, method_name string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Exception during execution of method %v: %v\n", method_name, r)
		}
	}()

	obj := reflect.ValueOf(rcv.Receiver)
	// fmt.Println(obj.Kind())
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
		case reflect.Bool:
			decider, _ := utils.IntGenerator(-100, 100)
			args[i] = reflect.ValueOf(utils.BooleanGenerator(decider))
		case reflect.Float32:
			arg, _ := utils.Float32Generator(-100, 100)
			args[i] = reflect.ValueOf(arg)
		case reflect.Float64:
			arg, _ := utils.Float64Generator(-100, 100)
			args[i] = reflect.ValueOf(arg)
		case reflect.Int16:
			arg := utils.Int16Generator(-100, 100)
			args[i] = reflect.ValueOf(arg)
		case reflect.Int8:
			arg := utils.Int8Generator(-100, 100)
			args[i] = reflect.ValueOf(arg)
		case reflect.Int32:
			arg, _ := utils.Int32Generator(-100, 100)
			args[i] = reflect.ValueOf(arg)
		case reflect.Int64:
			arg, _ := utils.Int64Generator(-100, 100)
			args[i] = reflect.ValueOf(arg)
		case reflect.Uint:
			arg := utils.UintGenerator(0, 100)
			args[i] = reflect.ValueOf(arg)
		case reflect.Uint64:
			arg := utils.Uint64Generator(0, 100)
			args[i] = reflect.ValueOf(arg)
		case reflect.Uint32:
			arg := utils.Uint32Generator(0, 100)
			args[i] = reflect.ValueOf(arg)
		case reflect.Uint16:
			arg := utils.Uint16Generator(0, 100)
			args[i] = reflect.ValueOf(arg)
		case reflect.Uint8:
			arg := utils.Uint8Generator(0, 10)
			args[i] = reflect.ValueOf(arg)
		case reflect.Ptr:
			if reflect.ValueOf(arg_type).Elem().Kind() == reflect.Struct {
				rcv_name := arg_type.Elem().Name()
				// Searches for the executors receivers
				for _, exec_rcv := range exec.ReceiversList {
					if exec_rcv.Name == rcv_name && exec_rcv.IsStar {
						exec_rcv.SetReceiverValues(exec.ReceiversList)
						arg := exec_rcv.Receiver
						args[i] = reflect.ValueOf(arg)
					}
				}
			} else {
				// Criar ponteiros para todos os tipos (bruhhhhhhhhhh)
			}
		}
	}

	// Talvez eu precise de mais um atributo em executor para armazenar os métodos. 
	results := method.Call(args)

	fmt.Printf("Method %v [%v] successfully executed: %v\n", method_name, args, results)
	// for _, result := range results {
		// fmt.Println(result)
	// }
}

func (exec *Executor) SimpleExecution() {
	for _, fn := range exec.FunctionsList {
		if !fn.IsMethod {
			exec.SetFuncArgs(fn)
			reflect_args := utils.ArgToReflectValue(fn.Args, fn.HasVariadic, fn)

			exec.ExecuteFunc(fn, reflect_args)

			
		} else {
			for _, rcv := range exec.ReceiversList {
				fmt.Println(rcv)
			}
		}
	}
}

func (exec *Executor) SetFuncArgs(fn *functions.Function) {
	var args, list_value []any
	var value any
	var list_arg_flag bool

	for i, tp := range fn.ArgTypesString {
		tpe := reflect.ValueOf(tp).Kind().String()

		if tp == "float64" {
			list_arg_flag = false
			value, _ = utils.Float64Generator()
		} else if (tp == "float32") {
			list_arg_flag = false
			value, _ = utils.Float32Generator()
		} else if (tp == "int") {
			list_arg_flag = false
			value, _ = utils.IntGenerator()
		} else if (tp == "int64") {
			list_arg_flag = false
			value, _ = utils.Int64Generator()
		} else if (tp == "int32") {
			list_arg_flag = false
			value, _ = utils.Int32Generator()
		} else if (tp == "uint") {
			list_arg_flag = false
			value = utils.UintGenerator()
		} else if (tp == "uint64") {
			list_arg_flag = false
			value = utils.Uint64Generator()
		} else if (tp == "uint32" || tpe == "rune") {
			list_arg_flag = false
			value = utils.Uint32Generator()
		} else if(tp == "uint16") {
			list_arg_flag = false
			value = utils.Uint16Generator()
		} else if (tp == "uint8" || tpe == "byte") {
			list_arg_flag = false
			value = utils.Uint8Generator()
		} else if (tp == "string") {
			list_arg_flag = false
			lenght, _ := utils.IntGenerator(1, 10)
			value = utils.StringGenerator(lenght)
		} else if (tp == "bool") {
			list_arg_flag = false
			decider, _ := utils.IntGenerator(0, 10)
			value = utils.BooleanGenerator(decider)
		} else if (tp == "[]string") { 					// Implementar outros tipos de listas
			list_arg_flag = true
			list_value = make([]any, 0)
			lenght, _ := utils.IntGenerator(0, 100)
				
			for j := 0; j < lenght; j++ {
				str_lenght, _ := utils.IntGenerator(0, 100)
				list_value = append(list_value, utils.StringGenerator(str_lenght))
			}
		} else if (tp == "[]byte") {
			list_arg_flag = true
			list_value = make([]any, 0)
			lenght, _ := utils.IntGenerator(0, 100)

			for j := 0; j < lenght; j++ {
				list_value = append(list_value, utils.Uint8Generator(0, 127))
			}
		} else if (tp == "io.Writer") {
			var buffer io.Writer = &bytes.Buffer{}
			value = buffer
		} else {
			/* Se o argumento é igual a uma struct */
			tpe := fn.ArgTypes[i].Kind().String()
			// fmt.Println(fn.ArgTypes[i].Kind().String())

			if tpe == "func" {
				var nil_value interface{} = nil 
				value = nil_value
			} else if tpe == "ptr" {
				data_kind := fn.ArgTypes[i].Elem().Kind()
				list_arg_flag = false

				if data_kind == reflect.Struct {
					struct_name := fn.ArgTypes[i].Elem().Name()

					// Será que não dá pra otimizar isso? utilizar um map ao invés de uma lista simples
					for _, rcv := range exec.ReceiversList {
						if rcv.Name == struct_name {
							rcv.SetReceiverValues(exec.ReceiversList)
							value = rcv.Receiver
						}
					}
				} 
			} else {
				list_arg_flag = false
				if tp == "error" {
					var nil_value interface{} = nil
					value = nil_value
				}
			}
		}

		if !list_arg_flag {
			fmt.Println(value)
			args = append(args, value)
		} else {
			args = append(args, list_value)
		}
	}

	fn.Args = args
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