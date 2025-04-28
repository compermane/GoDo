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

func (exec *Executor) PrintFunctions() {
	for _, fn := range exec.FunctionsList {
		fmt.Println(fn.Name)
	}
}

/* Sets function arguments. Its values should be the same througout the whole execution of GODO.
 * :param fn: Function to have arguments setted
 * :returns: None
 */
 func SetFuncArgs(fn *functions.Function, rcvs []*receiver.Receiver, 
	              seq *sequence.Sequence, global_ret_values map[string][]reflect.Value) []any {
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
			value, _ = utils.IntGenerator(-4096, 4096)
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
							value = nil
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
							value = nil
						}
					}
						// for _, rcv := range rcvs {
						// 	if rcv.Name == struct_name {
						// 		rcv.SetReceiverValues(rcvs)
						// 		value = CloneValue(rcv.Receiver)
						// 	}
						// }
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