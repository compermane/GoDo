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

func (rcv *Receiver) SetReceiverValues() {
	v := reflect.ValueOf(rcv.Receiver)

	if !rcv.IsStar {
		v = v.Addr()
	}

	v = v.Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		var value any

		switch rcv.AttrTypes[i].String() {
		case "float64":
			value, _ = utils.Float64Generator()
		case "float32":
			value, _ = utils.Float32Generator()
		case "int":
			value, _ = utils.IntGenerator()
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
		}

		rcv.AttrValues = append(rcv.AttrValues, value)
		field.Set(reflect.ValueOf(value))
	}

}

func get_methods_info(struct_type reflect.Type) []*functions.Function {
	mts := make([]*functions.Function, 0)

	for i := 0; i < struct_type.NumMethod(); i++ {
		method := struct_type.Method(i)
		method_type := method.Type

		args := make([]string, 0)
		returns := make([]string, 0)
		method_name := method.Name

		for j := 1; j < method_type.NumIn(); j++{
			args = append(args, method_type.In(j).String())
		}

		for k := 0; k < method_type.NumOut(); k++ {
			returns = append(returns, method_type.Out(k).String())
		}

		mts = append(mts, functions.InitFunction(method_name, struct_type.Name(), args, returns))
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

	return InitReceiver(rcv, struct_type.Name(), methods, method_names, attr_names, attr_types, true)
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


func SetAttrValues(rcv *Receiver) {
	var value interface{}
	var values []interface{}

	for _, attr := range rcv.AttrTypes {
		attr := attr.String()
		if attr == "float64" {
			value, _ = utils.Float64Generator()
		} else if (attr == "float32") {
			value, _ = utils.Float32Generator()
		} else if (attr == "int64") {
			value, _ = utils.Int64Generator()
		} else if (attr == "int32") {
			value, _ = utils.Int32Generator()
		}
		values = append(values, value)
	}

	rcv.AttrValues = values
}
