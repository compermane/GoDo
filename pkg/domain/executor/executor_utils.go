package executor

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"time"

	"github.com/compermane/ic-go/pkg/domain/functions"
	"github.com/compermane/ic-go/pkg/domain/receiver"
	"github.com/compermane/ic-go/pkg/domain/sequence"
	"github.com/compermane/ic-go/pkg/utils"
)

type DebugOpts struct {
	Dump               bool
	Debug              bool
	UseSequenceHashMap bool
	Iteration          int
}

func (exec *Executor) PrintFunctions() {
	for _, fn := range exec.FunctionsList {
		fmt.Println(fn.Name)
	}
}

func (exec *Executor) makeMapOfFunctions() map[string]bool {
	tested_function_map := make(map[string]bool, 0)

	for _, fn := range exec.FunctionsList {
		tested_function_map[fn.Name] = false	
	}

	return tested_function_map
}

/* Sets function arguments. Its values should be the same througout the whole execution of GODO.
 * :param fn: Function to have arguments setted
 * :param seq: Sequence where fn is going to be extended
 * :param global_ret_values: Global structs returned from the executor
 * :param create_structs: Creates a sample struct if true, otherwise sets nil
 * :returns: None
 */
 func SetFuncArgs(fn *functions.Function, seq *sequence.Sequence, rcvs []*receiver.Receiver,
	              global_ret_values map[string][]reflect.Value, create_structs bool) []any {
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
			value, _ = utils.IntGenerator(-1, 2)
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
			lenght, _ := utils.IntGenerator(1, 5)
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

			if tpe == "func" {
				var nil_value any = nil 
				value = nil_value
			} else if tpe == "ptr" {
				data_kind := fn.ArgTypes[i].Elem().Kind()
				list_arg_flag = false

				if data_kind == reflect.Struct {
					struct_name := fn.ArgTypes[i].String()

					if seq != nil {
						decider, _ := utils.IntGenerator(0, 2)
						switch decider {
						// Select a value v from a sequence that is already in the analysed sequence
						case 0:
							reflect_value, ok := seq.GetRandomReturnedValue(struct_name)
							if ok {
								value          = reflect_value
							} else {
								value          = nil
							}
						// Select a value v that is already returned from the nonErrorSeqs
						case 1:
							values     := global_ret_values[struct_name]
							
							if len(values) == 0 {
								value = nil 
								break
							}
							
							rand.Seed(time.Now().UnixNano())
							rand_index    := rand.Intn(len(values))
							
							reflect_value := values[rand_index]
							value          = reflect_value
						// Set a nil value
						case 2:
							if create_structs {
								for _, rcv := range rcvs {
									if struct_name == rcv.Name {
										rcv.SetReceiverValues(rcvs)
										value = CloneValue(rcv.Receiver)
										break
									}
								}
							} else {
								value = nil
							}
						}
					} else {
						decider, _ := utils.IntGenerator(1, 2)

						switch decider {
						case 1:
							values     := global_ret_values[struct_name]

							if len(values) == 0 {
								value = nil 
								break
							}

							rand.Seed(time.Now().UnixNano())
							rand_index    := rand.Intn(len(values))

							reflect_value := values[rand_index]
							value = reflect_value
						// Set a nil value
						case 2:
							if create_structs {
								for _, rcv := range rcvs {
									if struct_name == rcv.Name {
										rcv.SetReceiverValues(rcvs)
										value = CloneValue(rcv.Receiver)
									}
								}
							} else {
								value = nil
							}
						}
					}
				} 
			} else {
				list_arg_flag = false
				if tp == "error" {
					var nil_value any = nil
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

func CloneValue(original any) any {
	origVal := reflect.ValueOf(original)

	if origVal.Kind() == reflect.Ptr {
		origVal = origVal.Elem()
	}

	copyVal := reflect.New(origVal.Type()).Elem()

	copyVal.Set(origVal)

	return copyVal.Addr().Interface()
}

func UnwrapValue(val reflect.Value) any {
	// Se for um ponteiro, desreferencia
	if val.Kind() == reflect.Ptr {
		if !val.IsNil() {
			return val.Pointer() // Retorna o valor real da struct apontada
		}
		return nil // Se o ponteiro for nil, retorna nil
	}
	// Se já for uma struct, retorna diretamente a struct
	return val.Interface()
}

func (exec *Executor) AppendGlobalStruct(value_type string, value reflect.Value) {
	_, exist := exec.GlobalReceivers[value_type] 

	if exist {
		exec.GlobalReceivers[value_type] = append(exec.GlobalReceivers[value_type], value)
	} else {
		exec.GlobalReceivers[value_type] = make([]reflect.Value, 0)
		exec.GlobalReceivers[value_type] = append(exec.GlobalReceivers[value_type], value)
	}
}

func updateUntestedFunctions(seqs []*sequence.Sequence, func_map map[string]bool) {
	for _, seq := range seqs {
		if seq == nil {
			continue
		}
		for _, fn := range seq.Functions {
			func_map[fn.Name] = true
		}
	}
}

func updateErrorFunctions(seqs []*sequence.Sequence, func_map map[string]bool) {
	for _, seq := range seqs {
		for _, fn := range seq.Functions {
			if fn.HasError {
				func_map[fn.Name] = true
			} else {
				if func_map[fn.Name] {
					func_map[fn.Name] = false
				}
			}
		}
	}
}

func getUntestedFuncs(func_map map[string]bool) string {
	str_untested := ""

	for key, value := range func_map {
		if !value {
			str_untested += key + "\n"
		}
	}

	return str_untested
}

func getErrorFuncs(func_map map[string]bool) string {
	str_error := ""

	for key, value := range func_map {
		if value {
			str_error += key + "\n"
		}
	}

	return str_error
}

func makeMap(seqs []*sequence.Sequence) map[string]string {
	map_of_appearences := make(map[string]string)

	for _, seq := range seqs {
		if seq == nil {
			continue
		}
		for _, fn := range seq.Functions {
			if _, ok := map_of_appearences[fn.Name]; ok {
				if fn.HasError {
					map_of_appearences[fn.Name] += fmt.Sprintf("[%v] ", fn.Error)
				} else {
					map_of_appearences[fn.Name] += ""
				}
			} else {
				if fn.HasError {
					map_of_appearences[fn.Name] += fmt.Sprintf("[%v] ", fn.Error)
				} else {
					map_of_appearences[fn.Name] = ""
				}
			}
		}
	}

	return map_of_appearences
} 