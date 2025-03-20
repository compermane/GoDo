package executor

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"time"

	functions "github.com/compermane/ic-go/pkg/domain/functions"
	"github.com/compermane/ic-go/pkg/domain/receiver"
	"github.com/compermane/ic-go/pkg/domain/sequence"
	"github.com/compermane/ic-go/pkg/domain/testfunction"
	"github.com/compermane/ic-go/pkg/utils"
)

type Executor struct {
	Sequences           []*sequence.Sequence
	ReceiversList		[]*receiver.Receiver
	FunctionsList		[]*functions.Function
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
							  errorSequences []*sequence.Sequence) ([]*sequence.Sequence, []*sequence.Sequence) {
	fn  		  := functions.ChooseRandom(exec.FunctionsList)
	seq           := sequence.ChooseRandom(nonErrorSequences)
	args 		  := exec.SetFuncArgs(fn)
	reflect_args  := utils.ArgToReflectValue(args, fn.HasVariadic, fn)
	test_function := testfunction.NewTestFunction(fn, reflect_args)
	
	fmt.Printf("non error sequences len: %v\n", len(nonErrorSequences))
	if seq == nil {
		ok := exec.ExecuteTestFunc(test_function, reflect_args)

		if ok {
			new_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function}, nil)

			if nonErrorSequences != nil {
				if !sequence.VerifyExistence(nonErrorSequences, new_seq) {
					nonErrorSequences = append(nonErrorSequences, new_seq)
				}
			}
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
	} else {
		ok := exec.ExecuteTestFunc(test_function, reflect_args)

		if ok {
			new_seq := sequence.NewSequence([]*testfunction.TestFunction{test_function}, nil)
			seq.AppendSequence(new_seq)

			if nonErrorSequences != nil {
				if !sequence.VerifyExistence(nonErrorSequences, new_seq) {
					nonErrorSequences = append(nonErrorSequences, new_seq)
				}
			}
		}
	}

	return nonErrorSequences, errorSequences
}

/* Tries to execute sequences of functions. If its not possible due to some error, then the analysed sequences
 * are not concatenated. An object Executor is necessary for accessing its list of sequences.
 * :returns: None
 */
// func (exec *Executor) ExecuteSequences() {
// 	if len(exec.Sequences) == 0 {
// 		exec.SimpleExecution()
// 	}
	
// 	seq_1 := sequence.ChooseRandom(exec.Sequences)
// 	seq_2 := sequence.ChooseRandom(exec.Sequences)

// 	var ant_fn *functions.Function
// 	var can_append bool = true
// 	for i, fn := range seq_1.Functions {
// 		if i != 0 {
// 			for _, ret_value := range exec.FunctionsReturns[ant_fn] {
// 				fn.OverwriteArgValue(ret_value)
// 			}
// 		}

// 		if exec.ExecuteFunc(fn, fn.ReflectArgs) {
// 			ant_fn = fn
// 		}
// 	}

// 	for _, fn := range seq_2.Functions {
// 		for _, ret_value := range exec.FunctionsReturns[ant_fn] {
// 			fn.OverwriteArgValue(ret_value)
// 		}

// 		if exec.ExecuteFunc(fn, fn.ReflectArgs) {
// 			ant_fn = fn
// 		} else {
// 			can_append = false
// 			break
// 		}
// 	}

// 	if can_append {
// 		seq_1.AppendSequence(seq_2)
// 		if ok := sequence.VerifyExistence(exec.Sequences, seq_1); !ok {
// 			exec.Sequences = append(exec.Sequences, seq_1)
// 		}
// 	}
// }

/* Executes all executor functions only once.
 * :returns: None
 */
// func (exec *Executor) SimpleExecution() {
// 	for _, fn := range exec.FunctionsList {
// 		if !fn.IsMethod {
// 			args := exec.SetFuncArgs(fn)
// 			reflect_args := utils.ArgToReflectValue(args, fn.HasVariadic, fn)

// 			test_func := testfunction.NewTestFunction(fn, reflect_args)
// 			if exec.ExecuteFunc(fn, reflect_args) {
// 				exec.Sequences = append(exec.Sequences, sequence.NewSequence([]*functions.Function{fn}, nil))
// 			}

// 		} else {
// 			for _, rcv := range exec.ReceiversList {
// 				fmt.Println(rcv)
// 			}
// 		}
// 	}
// }

/* Sets function arguments. Its values should be the same througout the whole execution of GODO.
 * :param fn: Function to have arguments setted
 * :returns: None
 */
func (exec *Executor) SetFuncArgs(fn *functions.Function) []any {
	var args, list_value []any
	var value any
	var list_arg_flag bool

	for i, tp := range fn.ArgTypesString {
		tpe := fn.ArgTypes[i].Kind().String()
		if tp == "float64" {
			list_arg_flag = false
			value, _ = utils.Float64Generator()
		} else if (tp == "float32") {
			list_arg_flag = false
			value, _ = utils.Float32Generator()
		} else if (tp == "int" || tpe == "int") {
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
		} else if (tp == "interface {}") {
			decider, _ := utils.IntGenerator(1, 5)

			if decider == 1 {
				value, _ = utils.IntGenerator(-100, 100)
			} else if decider == 2 {
				value, _ = utils.Float32Generator(-100, 100)
			} else if decider == 3 {
				value, _ = utils.Float64Generator(-100, 100)
			} else if decider == 4 {
				lenght, _ := utils.IntGenerator(1, 100)
				value = utils.StringGenerator(lenght)
			} else if decider == 5 {
				decider, _ = utils.IntGenerator(0, 1)
				value = utils.BooleanGenerator(decider)
			}
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
			args = append(args, value)
		} else {
			args = append(args, list_value)
		}
	}

	return args
}

/* Main entry for GODO algorithm.
 * :param fns: List of functions
 * :param rcvs: List of receivers
 * :param algorithm: Algorithm for GODO execution
 * :param no_runs: Number of iterations of the chosen algorithm
 * :param timeout: Time limit for the chosen algorithm execution. Ignored if no_runs != 0
 * :returns: Pointer to an executor
 */
func ExecuteFuncs(fns, rcvs []any, algorithm string, no_runs, timeout int) *Executor {
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
				nonErrorSequences, errorSequences = exec.Randoop(nonErrorSequences, errorSequences)
			}
			fmt.Printf("final len: %v\n", len(nonErrorSequences))
			for _, seq := range nonErrorSequences {
				if seq == nil {
					continue
				}
				for _, fn := range seq.Functions {
					fmt.Printf("%v ",fn.Name)
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
					exec.Randoop(nonErrorSequences, errorSequences)
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