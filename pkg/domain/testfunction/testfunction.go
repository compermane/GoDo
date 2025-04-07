package testfunction

import (
	"math/rand"
	"reflect"
	"time"

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

func (tf *TestFunction) SelectRandomArg() string {
	if len(tf.ArgTypes) == 0 {
		return ""
	}
	
	source := rand.NewSource(time.Now().UnixNano())
	rng	   := rand.New(source)

	random_index := rng.Intn(len(tf.ArgTypes))

	return tf.ArgTypes[random_index].String()
}