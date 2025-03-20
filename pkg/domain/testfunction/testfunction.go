package testfunction

import (
	"reflect"

	"github.com/compermane/ic-go/pkg/domain/functions"
)

type TestFunction struct {
	*functions.Function
	ArgValues			[]reflect.Value
	RetValues			[]reflect.Value
	HasError			bool
	Error				any
}

func NewTestFunction(fn *functions.Function, args []reflect.Value) *TestFunction {
	return &TestFunction{
		fn,
		args,
		nil,
		false,
		nil,
	}
}