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
	lst  := make([]*functions.Function, 0)
	rcvs := make([]*receiver.Receiver, 0)
	glbl := make(map[string][]reflect.Value, 0)

	// mtds := receiver.GetMethodsFromReceivers(rcvs_lst)
	for _, fn := range fn_lst {
		lst = append(lst, functions.GetFunction(fn))
	}

	// for _, mtd := range mtds {
	// 	lst = append(lst, functions.GetFunction(mtd))
	// }
	for _, rcv := range rcvs_lst {
		rcvs = append(rcvs, receiver.GetReceiver(rcv))
	}

	return &Executor{
		Sequences: make([]*sequence.Sequence, 0),
		ReceiversList: rcvs,
		FunctionsList: lst,
		GlobalReceivers: glbl,
	}
}

func (exec *Executor) ExecuteTestFunc(fn *testfunction.TestFunction, args []reflect.Value, timeout time.Duration) (ok bool){
	result_chan := make(chan bool)
	
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Exception during execution of %v: %v\n", fn.Name, r)
				fn.HasError = true
				fn.Error = r

				result_chan <- false
				return 
			}
			result_chan <- true
		}()
		fn.RetValues = fn.Signature.Call(args)
	}()
	
	select {
	case ok = <-result_chan:
		return ok
	case <-time.After(timeout):
		fmt.Printf("Function %v timedout after %v\n", fn.Name, timeout)
		return false
	}		
}

func (exec *Executor) Randoop(nonErrorSequences []*sequence.Sequence, 
							  errorSequences 	[]*sequence.Sequence,
							  rcvs 				[]*receiver.Receiver,
							  debug 			bool,
							  timeout           time.Duration) ([]*sequence.Sequence, []*sequence.Sequence) {
	error_type    := reflect.TypeOf((*error)(nil)).Elem()
	fn  		  := functions.ChooseRandom(exec.FunctionsList)
	seq           := sequence.ChooseRandom(nonErrorSequences)
	args 		  := SetFuncArgs(fn, rcvs, seq, exec.GlobalReceivers)
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

	// Tratar já existência da sequência a ser formada
	if seq == nil {
		verify_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function}, nil)
		if sequence.VerifyExistence(nonErrorSequences, verify_seq) || sequence.VerifyExistence(errorSequences, verify_seq) {
			return nonErrorSequences, errorSequences
		}

	} else {
		verify_seq := seq.AppendSequence(sequence.NewSequence([]*testfunction.TestFunction{test_function}, nil))
		if sequence.VerifyExistence(nonErrorSequences, verify_seq) || sequence.VerifyExistence(errorSequences, verify_seq) {
			return nonErrorSequences, errorSequences
		}
	}

	// Caso a sequência selecionada é a vazia
	if seq == nil {
		ok 			     := exec.ExecuteTestFunc(test_function, reflect_args, timeout)
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
				nonErrorSequences = append(nonErrorSequences, new_seq)

				for _, ret_value := range test_function.RetValues {
					return_type := ret_value.Type().String()
					new_seq.AppendReturnedValue(return_type, ret_value)
				}
			}
		// Caso contrário, faz o append na lista de funções error
		} else {
			new_err_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function}, nil)
			errorSequences = append(errorSequences, new_err_seq)
		}
	// Caso contrário cria uma nova sequência unitária caso não haja erros
	// Faz o append em non error caso a nova sequência já não exista
	} else {
		ok               := exec.ExecuteTestFunc(test_function, reflect_args, timeout)
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
				nonErrorSequences = append(nonErrorSequences, new_seq)

				for _, ret_value := range test_function.RetValues {
					return_type := ret_value.Type().String()
					new_seq.ApplyExtensibleFlags(return_type, ret_value)
					new_seq.AppendReturnedValue(return_type, ret_value)

					if ret_value.Kind() == reflect.Struct || ret_value.Kind() == reflect.Pointer {
						exec.AppendGlobalStruct(return_type, ret_value)
					}
				}
			}
		} else {
			new_err_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function}, nil)
			new_err_seq  = seq.AppendSequence(new_err_seq)

			if errorSequences != nil {
				errorSequences = append(errorSequences, new_err_seq)
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
 * :param duration: Time limit for the chosen algorithm execution. Ignored if no_runs != 0
 * :param timmeout: Time limit for the execution of a single function
 * :param debug: Show information on the terminal
 * :param dump: Dump collected information on files
 * :returns: Pointer to an executor
 */
func ExecuteFuncs(fns, rcvs []any, algorithm string, no_runs, duration, timeout int, debug bool, dump bool) *Executor {
	exec := InitExecutor(fns, rcvs)
	func_timeout := time.Duration(timeout) * time.Second

	switch algorithm {
	case "baseline1":
		if no_runs > 0 {
			for i := 0; i < no_runs; i++ {
				// exec.SimpleExecution()
			}
		} else if duration > 0 {
			duration := time.Duration(duration) * time.Second
			timer := time.After(duration)
			
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
		} else if duration > 0 {
			duration := time.Duration(duration) * time.Second
			timer := time.After(duration)
			
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
				nonErrorSequences, errorSequences = exec.Randoop(nonErrorSequences, errorSequences, exec.ReceiversList, debug, time.Duration(func_timeout))
			}
			if dump {
				non_error := "-------------------Non error sequences-------------------\n"
				err_seq   := "-------------------Error sequences-----------------------\n"

				for i, seq := range nonErrorSequences {
					if seq != nil {
						non_error += fmt.Sprintf("Sequence %v: %v\n", i, seq.String())
					}
				}
				for i, seq := range errorSequences {
					if seq != nil {
						err_seq += fmt.Sprintf("Sequence %v: %v\n", i, seq.String())
					}
				}
				untested, err := getUntestedFuncs(nonErrorSequences, errorSequences, exec.FunctionsList)

				utils.DumpToFile("untested_functions.txt", untested)
				utils.DumpToFile("error_functions.txt", err)
				utils.DumpToFile("error_sequences.txt", err_seq)
				utils.DumpToFile("non_error_sequences.txt", non_error)
			} else {
				for i, seq := range nonErrorSequences {
					if seq != nil {
						fmt.Printf("Sequence %v: %v\n", i, seq.String())
					}
				}
			}
			fmt.Println("-----------Generated Objects----------")
			for struct_type, values := range exec.GlobalReceivers {
				fmt.Printf("%v: ", struct_type)
				for _, value := range values {
					fmt.Printf("%+v ", value)
				}
				fmt.Println()
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
					exec.Randoop(nonErrorSequences, errorSequences, exec.ReceiversList, debug, time.Duration(func_timeout))
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