package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// Generates random float64 values.
// :param params: Parameters for generations, where params[0] is the quantity of float values to be generated
// params[1] is the minimum value and params[2] is the maximum value. Defaults to 0 and 1 if not provided.
// :returns: A list containing the generated values and an error.
func Float64Generator(params ...float64) (value float64, e error) {
	min := 0.0
	max := 1.0

	if len(params) == 3 {
		min = params[0]
		max = params[1]
	}

	if max <= min {
		e = errors.New(fmt.Sprintf("Intervalo inválido: [%v - %v]", min, max))
		return 0, e
	}

	rand.Seed(time.Now().UnixNano())
	value = min + rand.Float64() * (max - min)

	return value, nil
}

// Generates random float32 values.
// :param params: Parameters for generations, where params[0] is the quantity of float values to be generated
// params[1] is the minimum value and params[2] is the maximum value. Defaults to 0 and 1 if not provided.
// :returns: A list containing the generated values and an error.
func Float32Generator(params ...float32) (value float32, e error) {
	min := float32(0)
	max := float32(1)

	if len(params) == 3 {
		min = params[0]
		max = params[1]
	}

	if max <= min {
		e = errors.New(fmt.Sprintf("Intervalo inválido: [%v - %v]", min, max))
		return 0, e
	}

	rand.Seed(time.Now().UnixNano())
	value = min + rand.Float32() * (max - min)
		
	return value, nil
}

func IntGenerator(params ...int) (value int, e error) {
	min := int(0)
	max := int(1)

	if len(params) == 2 {
		min = params[0]
		max = params[1]
	}

	if max <= min {
		e = errors.New(fmt.Sprintf("Intervalo inválido: [%v - %v]", min, max))
		return 0, e
	}

	rand.Seed(time.Now().UnixNano())
	value = min + rand.Intn(max-min)

	return value, nil
}

func Int64Generator(params ...int64) (value int64, e error) {
	min := int64(0)
	max := int64(1)

	if len(params) == 3 {
		min = params[0]
		max = params[1]
	}

	if max <= min {
		e = errors.New(fmt.Sprintf("Intervalo inválido: [%v - %v]", min, max))
		return 0, e
	}

	rand.Seed(time.Now().UnixNano())
	value = min + rand.Int63() * (max - min)


	return value, nil
}

func Int32Generator(params ...int32) (value int32, e error) {
	min := int32(0)
	max := int32(1)

	if len(params) == 3 {
		min = params[1]
		max = params[2]
	}

	if max <= min {
		e = errors.New(fmt.Sprintf("Intervalo inválido: [%v - %v]", min, max))
		return 0, e
	}

	rand.Seed(time.Now().UnixNano())
	value = min + rand.Int31() * (max - min)

	return value, nil
}

func StringGenerator(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	result := make([]byte, length) 
	for  i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}

	return string(result)
}

func BooleanGenerator(decider int) bool {
	return decider % 2 == 0
}

func Float64ToReflectValues(args []float64) (values []reflect.Value) {
	for _, arg := range args {
		value := reflect.ValueOf(arg)
		values = append(values, value)
	}

	return values
}