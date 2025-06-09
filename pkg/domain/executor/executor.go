package executor

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	functions "github.com/compermane/ic-go/pkg/domain/functions"
	"github.com/compermane/ic-go/pkg/domain/receiver"
	"github.com/compermane/ic-go/pkg/domain/sequence"
	"github.com/compermane/ic-go/pkg/domain/testfunction"
	"github.com/compermane/ic-go/pkg/utils"
)

type Executor struct {
	Sequences           []*sequence.Sequence			// Sequências geradas pelo algoritmo 
	FunctionsList		[]*functions.Function 			// Lista de funções para serem testadas
	ReceiversList		[]*receiver.Receiver			// Lista de receivers para serem testados
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

	for _, rcv := range rcvs_lst {
		rcvs = append(rcvs, receiver.GetReceiver(rcv))
	}
	
	return &Executor{
		Sequences: make([]*sequence.Sequence, 0),
		FunctionsList: lst,
		ReceiversList: rcvs,
		GlobalReceivers: glbl,
	}
}

func (exec *Executor) ExecuteTestFunc(fn *testfunction.TestFunction, args []reflect.Value, timeout time.Duration) (ok bool){
	result_chan := make(chan bool)
	
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// fmt.Printf("Exception during execution of %v: %v\n", fn.Name, r)
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
		if ok {
			fn.HasError = false
		}
		return ok
	case <-time.After(timeout):
		fmt.Printf("Function %v timedout after %v\n", fn.Name, timeout)
		return false
	}		
}

/* Entry for Randoop's algorithm.
 * :param nonErrorSequences: List of non error sequences. Should initially be empty, with only the nil sequence.
 * :param errorSequences: List of error sequences. Should initially be empty.
 * :param debug: Flag for displaying the algorithm information on the terminal.
 * :param timeout: Time in seconds for function execution timeout. A sequence is put on errorSequences if the timeout is violated.
 * :returns: list of resulting error sequences and non error sequences for algorithm interation.
 */
func (exec *Executor) Randoop(nonErrorSequences   []*sequence.Sequence, 
							  errorSequences 	  []*sequence.Sequence,
							  debug 		 	  bool,
							  create_structs      bool,
							  timeout             time.Duration,
							  error_sequences_map map[uint64]bool) ([]*sequence.Sequence, []*sequence.Sequence) {
	error_type    := reflect.TypeOf((*error)(nil)).Elem()
	fn  		  := functions.ChooseRandom(exec.FunctionsList)
	seq           := sequence.ChooseRandom(nonErrorSequences)
	args 		  := SetFuncArgs(fn, seq, exec.ReceiversList, exec.GlobalReceivers, create_structs)
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
		verify_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function})
		if error_sequences_map != nil {
			if sequence.VerifyExistence(nonErrorSequences, verify_seq) || sequence.VerifyExistence(errorSequences, verify_seq) {
				return nonErrorSequences, errorSequences
			}
		} else {
			// Aqui, verificar existência em nonErrorSequences pode ser otimizado
			if sequence.VerifyExistenceByHash(error_sequences_map, verify_seq.SequenceID) || sequence.VerifyExistence(nonErrorSequences, verify_seq) {
				return nonErrorSequences, errorSequences
			}
		}

	} else {
		verify_seq := seq.AppendSequence(sequence.NewSequence([]*testfunction.TestFunction{test_function}))
		if error_sequences_map != nil {
			if sequence.VerifyExistence(nonErrorSequences, verify_seq) || sequence.VerifyExistence(errorSequences, verify_seq) {
				return nonErrorSequences, errorSequences
			}
		} else {
			if sequence.VerifyExistence(nonErrorSequences, verify_seq) || sequence.VerifyExistenceByHash(error_sequences_map, verify_seq.SequenceID) {
				return nonErrorSequences, errorSequences
			}
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
			new_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function})

			if nonErrorSequences != nil {
				nonErrorSequences = append(nonErrorSequences, new_seq)

				for _, ret_value := range test_function.RetValues {
					return_type := ret_value.Type().String()
					new_seq.AppendReturnedValue(return_type, ret_value)
				}
			}
		// Caso contrário, faz o append na lista de funções error
		} else {
			new_err_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function})
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
			new_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function})
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
			new_err_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function})
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
 *
 * :param fns: List of functions
 * :param rcvs: List of receivers
 * :param algorithm: Algorithm for GODO execution
 * :param no_runs: Number of iterations of the chosen algorithm
 * :param duration: Time limit for the chosen algorithm execution (seconds). Ignored if no_runs != 0
 * :param timmeout: Time limit for the execution of a single function (seconds)
 * :param debug: Show information on the terminal
 * :param dump: Dump collected information on files
 * :returns: Pointer to an executor
 *
 */
func ExecuteFuncs(fns, rcvs []any, algorithm string, no_runs, duration, timeout int, debugOpts DebugOpts) *Executor {
	exec := InitExecutor(fns, rcvs)
	func_timeout := time.Duration(timeout) * time.Second

	var errorSequencesHashMap map[uint64]bool = nil
	var untestedFuctions      map[string]bool = exec.makeMapOfFunctions()
	var errorFunctions       map[string]bool = exec.makeMapOfFunctions()

	if debugOpts.UseSequenceHashMap {
		errorSequencesHashMap = make(map[uint64]bool, 0)
	}

	switch algorithm {
	case "feedback_directed":
		nonErrorSequences     := make([]*sequence.Sequence, 0)
		errorSequences        := make([]*sequence.Sequence, 0)

		nonErrorSequences = append(nonErrorSequences, nil)
		if no_runs > 0 {
			for i := 0; i < no_runs; i++ {
				nonErrorSequences, errorSequences = exec.Randoop(nonErrorSequences, errorSequences, debugOpts.Debug, false, time.Duration(func_timeout), errorSequencesHashMap)
				runtime.GC()
			}
			if debugOpts.Dump {
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
				updateUntestedFunctions(errorSequences, untestedFuctions)
				updateUntestedFunctions(nonErrorSequences, untestedFuctions)
				updateErrorFunctions(errorSequences, errorFunctions)
				untested := getUntestedFuncs(untestedFuctions)
				err      := getErrorFuncs(errorFunctions)

				utils.DumpToFile(fmt.Sprintf("untested_functions-%v.txt", debugOpts.Iteration), untested)
				utils.DumpToFile(fmt.Sprintf("error_functions-%v.txt", debugOpts.Iteration), err)
				utils.DumpToFile(fmt.Sprintf("error_sequences-%v.txt", debugOpts.Iteration), err_seq)
				utils.DumpToFile(fmt.Sprintf("non_error_sequences-%v.txt", debugOpts.Iteration), non_error)
			} else {
				for i, seq := range nonErrorSequences {
					if seq != nil {
						fmt.Printf("Sequence %v: %v\n", i, seq.String())
					}
				}
			}
		} else if duration > 0 {
			duration := time.Duration(duration) * time.Second
			timer := time.After(duration)
			
			n := 0
			for {
				select {
				case <-timer:
					fmt.Printf("Executed %v funcs\n\n", n)

					if debugOpts.Dump {
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
						updateUntestedFunctions(errorSequences, untestedFuctions)
						updateUntestedFunctions(nonErrorSequences, untestedFuctions)
						updateErrorFunctions(errorSequences, errorFunctions)

						untested := getUntestedFuncs(untestedFuctions)
						err      := getErrorFuncs(errorFunctions)
						utils.DumpToFile(fmt.Sprintf("untested_functions-%v.txt", debugOpts.Iteration), untested)
						utils.DumpToFile(fmt.Sprintf("error_functions-%v.txt", debugOpts.Iteration), err)
						utils.DumpToFile(fmt.Sprintf("error_sequences-%v.txt", debugOpts.Iteration), err_seq)
						utils.DumpToFile(fmt.Sprintf("non_error_sequences-%v.txt", debugOpts.Iteration), non_error)
					} else {
						for i, seq := range nonErrorSequences {
							if seq != nil {
								fmt.Printf("Sequence %v: %v\n", i, seq.String())
							}
						}
					}
					return exec
				default:
					nonErrorSequences, errorSequences = exec.Randoop(nonErrorSequences, errorSequences, debugOpts.Debug, false, time.Duration(func_timeout), errorSequencesHashMap)

					// Update hash map e untested functions
					if debugOpts.UseSequenceHashMap {
						errorSequencesHashMap = sequence.UpdateHashMap(errorSequencesHashMap, errorSequences)
						
						updateErrorFunctions(errorSequences, errorFunctions)
						updateUntestedFunctions(errorSequences, untestedFuctions)
						updateUntestedFunctions(nonErrorSequences, untestedFuctions)
					}

					runtime.GC()
					n = n + 1

					if len(errorSequences) % 2000 == 0 && debugOpts.UseSequenceHashMap {
						err_seq := ""
						for i, seq := range errorSequences {
							if seq != nil {
								err_seq += fmt.Sprintf("Sequence %v: %v\n", i, seq.String())
							}
						}	
						utils.DumpToFile(fmt.Sprintf("error_sequences-%v.txt", debugOpts.Iteration), err_seq)
						errorSequences = nil
						errorSequences = make([]*sequence.Sequence, 0)
					}
				}
			}
		} else {
			panic("Invalid number of runs or timeout duration\n")
		}
	case "feedback_directed_struct_generation":
		nonErrorSequences     := make([]*sequence.Sequence, 0)
		errorSequences        := make([]*sequence.Sequence, 0)
		
		nonErrorSequences = append(nonErrorSequences, nil)
		if no_runs > 0 {
			for i := 0; i < no_runs; i++ {
				nonErrorSequences, errorSequences = exec.Randoop(nonErrorSequences, errorSequences, debugOpts.Debug, true, time.Duration(func_timeout), errorSequencesHashMap)
				runtime.GC()
			}
			if debugOpts.Dump {
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
				updateUntestedFunctions(errorSequences, untestedFuctions)
				updateUntestedFunctions(nonErrorSequences, untestedFuctions)
				updateErrorFunctions(errorSequences, errorFunctions)
				untested := getUntestedFuncs(untestedFuctions)
				err      := getErrorFuncs(errorFunctions)

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

		} else if timeout > 0 {
			duration := time.Duration(duration) * time.Second
			timer := time.After(duration)
			
			n := 0
			for {
				select {
				case <-timer:
					fmt.Printf("Executed %v funcs\n\n", n)
					return exec
				default:
					exec.Randoop(nonErrorSequences, errorSequences, debugOpts.Debug, true, time.Duration(func_timeout), errorSequencesHashMap)
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