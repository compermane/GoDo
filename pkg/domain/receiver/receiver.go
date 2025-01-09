// Package for function receivers

package receiver

import (
	"reflect"

	"github.com/compermane/ic-go/pkg/domain/functions"
	"github.com/compermane/ic-go/pkg/utils"
)

type Receiver struct {
	Receiver    any
	Name		string
	Methods		[]*functions.Function
	MethodNames	[]string
	AttrNames	[]string
	AttrTypes	[]reflect.Type
	AttrValues	[]any
	IsStar		bool
}

func InitReceiver(rcv any, name string, methods []*functions.Function, method_names, attr_names []string, attr_types []reflect.Type, is_star bool) *Receiver {
	return &Receiver{
		Receiver: rcv,
		Name: name,
		Methods: methods,
		MethodNames: method_names,
		AttrNames: attr_names,
		AttrTypes: attr_types,
		IsStar: is_star,
	}
}

/* Sets attributes values for a given receiver via reflection.
 * :param rcvs_list: Receivers list declared for the main executor. Its particularly useful for when a struct has a struct attirbute. Can be nil.
 * :returns: None
 */
func (rcv *Receiver) SetReceiverValues(rcvs_list []*Receiver) {
	v := reflect.ValueOf(rcv.Receiver)

	// First, create a pointer to the struct (if its not already) -> this is because we want
	// to manipulate its fields values with reflection
	if !rcv.IsStar {
		ptr := reflect.New(v.Type())
		ptr.Elem().Set(v)
		v = ptr
	}

	v = v.Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		field_type := t.Field(i)
		// fmt.Println(field_type.Type.Kind())

		// fmt.Println(field.Kind())
		if !field_type.IsExported() {
			// fmt.Printf("Attribute %v not exportable. Continuing...\n", rcv.AttrNames[i])
			continue
		}

		var arr_type		string
		var value 			any
		var arr_value_flag 	bool = false
		var arr_value 		[]any
		// var map_value_flag	bool = false

		if rcv.AttrTypes[i].Kind().String() == "func" {
			value = nil
		} else if rcv.AttrTypes[i].Kind().String() == "struct" {
			rcv_name := rcv.AttrTypes[i].Name()

			for _, rcvs_list_name := range rcvs_list {
				if rcv_name == rcvs_list_name.Name {
					rcvs_list_name.SetReceiverValues(rcvs_list)
					break;
				}
			}
		} else {
			switch rcv.AttrTypes[i].String() {
			case "float64":
				value, _ = utils.Float64Generator(1, -100, 100)
			case "float32":
				value, _ = utils.Float32Generator()
			case "int":
				value, _ = utils.IntGenerator(-100, 100)
			case "int32":
				value, _ = utils.Int32Generator()
			case "int64":
				value, _ = utils.Int64Generator()
			case "bool":
				decider, _ := utils.IntGenerator(0, 10000)
				value = utils.BooleanGenerator(decider)
			case "string":
				lenght, _ := utils.IntGenerator(0, 100)
				value = utils.StringGenerator(lenght)
			case "uint":
				value = utils.UintGenerator(0, 100)
			case "uint64":
				value = utils.Uint64Generator(0, 100)
			case "uint32":
				value = utils.Uint32Generator(0, 100)
			case "uint16":
				value = utils.Uint16Generator(0, 100)
			case "uint8":
				value = utils.Uint8Generator(0, 100)
			case "[]string":
				arr_value_flag = true
				arr_type = "string"
				arr_value = make([]any, 0)

				lenght, _ := utils.IntGenerator(1, 100)
				for j := 0; j < lenght; j++ {
					str_lenght, _ := utils.IntGenerator(1, 100)
					arr_value = append(arr_value, utils.StringGenerator(str_lenght))
				}
			case "[]byte":
			case "[]uint8":
				arr_value_flag = true
				arr_value = make([]any, 0)

				lenght, _ := utils.IntGenerator(1, 100)
				for j := 0; j < lenght; j++ {
					arr_value = append(arr_value, utils.Uint8Generator(0, 127))
				}
			case "map[string]string":
				map_type := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(""))
				my_map := reflect.MakeMap(map_type)
				num_keys, _ := utils.IntGenerator(1, 10)
				
				for i := 0; i < num_keys; i++ {
					key_size, _ := utils.IntGenerator(1, 100)
					key := utils.StringGenerator(key_size)
					
					value_size, _ := utils.IntGenerator(1, 100)
					value := utils.StringGenerator(value_size)
					
					my_map.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
				}
				value = my_map.Interface()
			default:
				if rcv.AttrTypes[i].Kind() == reflect.Func {
					func_type := rcv.AttrTypes[i]
					arg_types_str := make([]string, 0)
					ret_types_str := make([]string, 0)

					for j := 0; j < func_type.NumIn(); j++ {
						arg_types_str = append(arg_types_str, func_type.In(j).String())
					}

					for j := 0; j < func_type.NumOut(); j++ {
						ret_types_str = append(ret_types_str, func_type.Out(j).String())
					}

				} else {
					panic("Unrecognized type: " + rcv.AttrTypes[i].String())
				}
			}

			if !arr_value_flag {
				rcv.AttrValues = append(rcv.AttrValues, value)
				field.Set(reflect.ValueOf(value))
			} else {
				if arr_type == "string" {
					// fmt.Println(arr_value...)
					rcv.AttrValues = append(rcv.AttrValues, utils.AnySliceToStringSlice(arr_value))
					field.Set(reflect.ValueOf(utils.AnySliceToStringSlice(arr_value)))
				}
			}
		}
	}
	// Copies the generated receiver and dereferences if it is not a pointer. 
	// This is because the pointer ceases to exist after this function execution (i think)
	if !rcv.IsStar {
		updated_receiver := v.Interface()
		rcv.Receiver = updated_receiver
	} 
}

func get_methods_info(struct_type reflect.Type) []*functions.Function {
	mts := make([]*functions.Function, 0)

	for i := 0; i < struct_type.NumMethod(); i++ {
		method := struct_type.Method(i)
		method_type := method.Type

		args := make([]reflect.Type, 0)
		args_str := make([]string, 0)
		returns := make([]string, 0)
		is_variadic := false
		method_name := method.Name
		
		for j := 1; j < method_type.NumIn(); j++{
			if j == method_type.NumIn() - 1 {
				if method_type.IsVariadic() {
					is_variadic = true
				}
			}
			args = append(args, method_type.In(j))
			args_str = append(args_str, method_type.In(j).String())
		}

		for k := 0; k < method_type.NumOut(); k++ {
			returns = append(returns, method_type.Out(k).String())
		}

		mts = append(mts, functions.InitFunction(method_name, struct_type.Name(), is_variadic, args, args_str, returns))
	}

	return mts
}

func proccess_receiver(struct_type reflect.Type, rcv any) *Receiver {
	methods := get_methods_info(struct_type)
	
	attr_names := make([]string, 0)
	attr_types := make([]reflect.Type, 0)
	method_names := make([]string, 0)

	for i := 0; i < struct_type.NumField(); i++ {
		attr_names = append(attr_names, struct_type.Field(i).Name)
		attr_types = append(attr_types, struct_type.Field(i).Type)
	}
	
	for i := 0; i < struct_type.NumMethod(); i++ {
		method_names = append(method_names, struct_type.Method(i).Name)
	}

	if struct_type.Kind() == reflect.Ptr {
    	ptr_type := reflect.PtrTo(struct_type)
    	for i := 0; i < ptr_type.NumMethod(); i++ {
        	method_names = append(method_names, ptr_type.Method(i).Name)
    	}
	}

	return InitReceiver(rcv, struct_type.Name(),  methods, method_names, attr_names, attr_types, false)
}


func process_star_receiver(struct_type reflect.Type, rcv any) *Receiver {
	methods := get_methods_info(struct_type)

	attr_names := make([]string, 0)
	attr_types := make([]reflect.Type, 0)
	method_names := make([]string, 0)

	for i := 0; i < struct_type.Elem().NumField(); i++ {
		attr_names = append(attr_names, struct_type.Elem().Field(i).Name)
		attr_types = append(attr_types, struct_type.Elem().Field(i).Type)
	}
	
	for i := 0; i < struct_type.NumMethod(); i++ {
		method_names = append(method_names, struct_type.Method(i).Name)
	}

	if struct_type.Kind() == reflect.Ptr {
    	ptr_type := reflect.PtrTo(struct_type)
    	for i := 0; i < ptr_type.NumMethod(); i++ {
        	method_names = append(method_names, ptr_type.Method(i).Name)
    	}
	}

	return InitReceiver(rcv, struct_type.Elem().Name(), methods, method_names, attr_names, attr_types, true)
}

func GetReceiver(rcv any) *Receiver {
	struct_type := reflect.TypeOf(rcv)

	switch struct_type.Kind() {
	case reflect.Struct:
		return proccess_receiver(struct_type, rcv)
	case reflect.Ptr:
		return process_star_receiver(struct_type, rcv)
	default:
		panic("Expected reflect.Struct or reflect.Ptr argument, received " + struct_type.Kind().String())
	}
}
