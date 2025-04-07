package executor

import (
	"fmt"
	"reflect"
	"time"

	functions "github.com/compermane/ic-go/pkg/domain/functions"
	"github.com/compermane/ic-go/pkg/domain/receiver"
	"github.com/compermane/ic-go/pkg/domain/sequence"
	"github.com/compermane/ic-go/pkg/domain/testfunction"
	"github.com/compermane/ic-go/pkg/utils"
)

type Executor struct {
	Sequences           []*sequence.Sequence			// Sequências geradas pelo algoritmo 
	ReceiversList		[]*receiver.Receiver			// Lista de receivers para execução de seus métodos/execução em que são argumento
	FunctionsList		[]*functions.Function 			// Lista de funções para serem testadas
	GlobalReceivers		map[string][]reflect.Value      // Mapa contendo todas os receivers retornados durante a execução do algoritmo
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
		Sequences: make([]*sequence.Sequence, 0),
		ReceiversList: rcvs,
		FunctionsList: lst,
	}
}

func (exec *Executor) ExecuteTestFunc(fn *testfunction.TestFunction, args []reflect.Value) (ok bool){
	ok = false
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Exception during execution of %v: %v\n", fn.Name, r)
			fn.HasError = true
			fn.Error = r
			ok = false
			} else {
				ok = true
			}
			}()

	time.Sleep(15 * time.Millisecond)
	fn.RetValues = fn.Signature.Call(args)
	// fmt.Printf("Function %v ( %v ) succesfully executed: %v\n", fn.Name, fn.Args, exec.GetFunctionReturns(fn))

	return 
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

func (exec *Executor) Randoop(nonErrorSequences []*sequence.Sequence, 
							  errorSequences 	[]*sequence.Sequence,
							  rcvs 				[]*receiver.Receiver,
							  debug 			bool) ([]*sequence.Sequence, []*sequence.Sequence) {
	error_type    := reflect.TypeOf((*error)(nil)).Elem()
	fn  		  := functions.ChooseRandom(exec.FunctionsList)
	seq           := sequence.ChooseRandom(nonErrorSequences)
	args 		  := SetFuncArgs(fn, rcvs)

	for _, arg := range args {
		fmt.Printf("%+v ", arg)
	}
	fmt.Println()
	reflect_args  := utils.ArgToReflectValue(args, fn.HasVariadic, fn)
	test_function := testfunction.NewTestFunction(fn, reflect_args)
	
	if debug {
		fmt.Printf("non error sequences len: %v\n", len(nonErrorSequences))
		if seq == nil {
			fmt.Println("selected sequence: nil")
		} else {
			fmt.Printf("selected sequence: %v\n", seq.String())
		}
		fmt.Printf("selected func: %v [ ", test_function.Name)
		for _, value := range reflect_args {
			fmt.Printf("%+v ", value)
		}
		fmt.Printf("]\n")
		for _, arg := range test_function.ArgValues {
			fmt.Printf("%+v ", arg)
		}
		fmt.Println()
	}

	// Caso a sequência selecionada é a vazia
	if seq == nil {
		ok 			     := exec.ExecuteTestFunc(test_function, reflect_args)
		no_error_returns := true

		for _, ret_val := range test_function.RetValues {
			if ret_val.Kind() == reflect.Ptr    || ret_val.Kind() == reflect.Interface ||
			   ret_val.Kind() == reflect.Slice  || ret_val.Kind() == reflect.Map       ||
			   ret_val.Kind() == reflect.Chan   || ret_val.Kind() == reflect.Func {
				if ret_val.IsNil() {
					continue
				}
			}
			if ret_val.Type().Implements(error_type) {
				no_error_returns = false
				break
			}
		}
		// Se a execução não deu panic, cria uma nova sequência unitária
		// Faz o append caso não a função não exista na lista de funções non error
		if ok && no_error_returns {
			new_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function}, nil)

			if nonErrorSequences != nil {
				if !sequence.VerifyExistence(nonErrorSequences, new_seq) {
					nonErrorSequences = append(nonErrorSequences, new_seq)

					for _, ret_value := range test_function.RetValues {
						return_type := ret_value.Type().String()
						new_seq.AppendReturnedValue(return_type, ret_value)
					}
				}
			}
		// Caso contrário, faz o append na lista de funções error
		} else {
			new_err_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function}, nil)
			if errorSequences != nil {
				if !sequence.VerifyExistence(errorSequences, new_err_seq) {
					errorSequences = append(errorSequences, new_err_seq)
				}
			} else {
				errorSequences = append(errorSequences, new_err_seq)
			}
		}
	// Caso contrário cria uma nova sequência unitária caso não haja erros
	// Faz o append em non error caso a nova sequência já não exista
	} else {
		decider, _ 				:= utils.IntGenerator(0, 1)
		select_random_ret_value := utils.BooleanGenerator(decider)

		if select_random_ret_value {
			arg_type := test_function.SelectRandomArg()
			if arg_type != "" {
				ret_val, ok := seq.GetRandomReturnedValue(arg_type)
				if ok {
					for i, arg := range reflect_args {
						if !arg.IsValid() {
							continue
						}
						if arg.Type() == ret_val.Type() {
							if debug {
								fmt.Printf("Overwriting %+v with %+v\n", arg, ret_val)
							}
							reflect_args[i] = ret_val
							break
						}
					}
				}
			}
		}
		ok               := exec.ExecuteTestFunc(test_function, reflect_args)
		no_error_returns := true
		for _, ret_val := range test_function.RetValues {
			if ret_val.Kind() == reflect.Ptr    || ret_val.Kind() == reflect.Interface ||
			   ret_val.Kind() == reflect.Slice  || ret_val.Kind() == reflect.Map       ||
			   ret_val.Kind() == reflect.Chan   || ret_val.Kind() == reflect.Func {
				if ret_val.IsNil() {
					continue
				}
			}
			if ret_val.Type().Implements(error_type) {
				no_error_returns = false
				break
			}
		}

		if ok && no_error_returns {
			new_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function}, nil)
			new_seq  = seq.AppendSequence(new_seq)

			if nonErrorSequences != nil {
				if !sequence.VerifyExistence(nonErrorSequences, new_seq) {
					nonErrorSequences = append(nonErrorSequences, new_seq)

					for _, ret_value := range test_function.RetValues {
						return_type := ret_value.Type().String()
						new_seq.AppendReturnedValue(return_type, ret_value)
					}
				}
			}
		} else {
			new_err_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function}, nil)
			new_err_seq  = seq.AppendSequence(new_err_seq)

			if errorSequences != nil {
				if !sequence.VerifyExistence(errorSequences, new_err_seq) {
					errorSequences = append(errorSequences, new_err_seq)
				}
			}
		}
	}

	if debug {
		fmt.Println("-----Non error sequences-----")
		for _, seq := range nonErrorSequences {
			if seq != nil {
				fmt.Println("-------------------")
				fmt.Printf("%v %v\n", seq.String(), seq.SequenceID)

				for chave, valores := range seq.ReturnedValues {
					fmt.Printf("%v: ", chave)
					for _, valor := range valores {
						fmt.Printf("%+v ", valor)
					}
					fmt.Println()
				}
			}
		}
		fmt.Println("Error sequences")
		for _, seq := range errorSequences {
			if seq != nil {
				fmt.Println(seq.String())
			}
		}
		fmt.Println("-----end of iteration-----")
	}

	return nonErrorSequences, errorSequences
}


/* Main entry for GODO algorithm.
 * :param fns: List of functions
 * :param rcvs: List of receivers
 * :param algorithm: Algorithm for GODO execution
 * :param no_runs: Number of iterations of the chosen algorithm
 * :param timeout: Time limit for the chosen algorithm execution. Ignored if no_runs != 0
 * :returns: Pointer to an executor
 */
func ExecuteFuncs(fns, rcvs []any, algorithm string, no_runs, timeout int, debug bool) *Executor {
	exec := InitExecutor(fns, rcvs)

	switch algorithm {
	case "baseline1":
		if no_runs > 0 {
			for i := 0; i < no_runs; i++ {
				// exec.SimpleExecution()
			}
		} else if timeout > 0 {
			timeout := time.Duration(timeout) * time.Second
			timer := time.After(timeout)
			
			n := 0
			for {
				select {
				case <-timer:
					fmt.Printf("Executed %v funcs\n\n", n)
					return exec
				default:
					// exec.SimpleExecution()
					n = n + 1
				}
			}
		} else {
			panic("Invalid number of runs or timeout duration\n")
		}
	case "baseline2":
		if no_runs > 0 {
			for i := 0; i < no_runs; i++ {
				fmt.Println(i)
				// exec.ExecuteSequences()
			}
		} else if timeout > 0 {
			timeout := time.Duration(timeout) * time.Second
			timer := time.After(timeout)
			
			n := 0
			for {
				select {
				case <-timer:
					fmt.Printf("Executed %v funcs\n\n", n)
					return exec
				default:
					// exec.SimpleExecution()
					n = n + 1
				}
			}
		} else {
			panic("Invalid number of runs or timeout duration\n")
		}
	case "feedback_directed":
		nonErrorSequences := make([]*sequence.Sequence, 0)
		errorSequences    := make([]*sequence.Sequence, 0)
		
		nonErrorSequences = append(nonErrorSequences, nil)
		if no_runs > 0 {
			for i := 0; i < no_runs; i++ {
				nonErrorSequences, errorSequences = exec.Randoop(nonErrorSequences, errorSequences, exec.ReceiversList, debug)
			}
			for i, seq := range nonErrorSequences {
				if seq != nil {
					fmt.Printf("Sequence %v: %v\n", i, seq.String())
				}
			}
		} else if timeout > 0 {
			timeout := time.Duration(timeout) * time.Second
			timer := time.After(timeout)
			
			n := 0
			for {
				select {
				case <-timer:
					fmt.Printf("Executed %v funcs\n\n", n)
					return exec
				default:
					exec.Randoop(nonErrorSequences, errorSequences, exec.ReceiversList, debug)
					n = n + 1
				}
			}

		} else {
			panic("Invalid number of runs or timeout duration\n")
		}
	default:
		panic("No algorithm named: " + algorithm + "\n")
	}

	return exec
}