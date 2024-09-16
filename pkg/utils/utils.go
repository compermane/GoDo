package utils

import (
	"math/rand"
	"reflect"
)

type Function struct {
	FuncName  	string
	ModName		string
}

func InputFactory(types []reflect.Type) ([]reflect.Value, error) {
	int64_type := reflect.TypeOf(int(0))
	float64_type := reflect.TypeOf(float64(0))

	var arg_values []reflect.Value

	for _, arg_type := range types {
		if arg_type == float64_type {
			arg_values = append(arg_values, reflect.ValueOf(rand.Float64()))
		} else if arg_type == int64_type {
			arg_values = append(arg_values, reflect.ValueOf(rand.Int63()))
		}
	}

	return arg_values, nil
}