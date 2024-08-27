package utils

import (
	"errors"
	"reflect"
)

type Function interface{}

func GetFunctionInformation(f Function) ([]reflect.Type, error) {
	func_type := reflect.TypeOf(f)

	if func_type.Kind() != reflect.Func {
		return nil, errors.New("O argumento f precisa ser uma função")
	}

	var arg_types []reflect.Type

	for i := 0; i < func_type.NumIn(); i++ {
		arg_type := func_type.In(i)
		arg_types = append(arg_types, arg_type)
	}

	return arg_types, nil
}