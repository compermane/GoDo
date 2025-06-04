package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
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
		min = params[1]
		max = params[2]
	}

	if max <= min {
		e = errors.New(fmt.Sprintf("Intervalo inválido: [%v - %v]", min, max))
		return 0, e
	}

	rand.Seed(time.Now().UnixNano())
	value = min + rand.Float64() * (max - min)

	time.Sleep(1 * time.Millisecond)

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
		
	time.Sleep(1 * time.Millisecond)

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
	value = min + rand.Intn(max-min + 1)

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

	time.Sleep(1 * time.Millisecond)

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

	time.Sleep(1 * time.Millisecond)

	return value, nil
}

func Int16Generator(params ...int) int16 {
	min := int16(0)
	max := int16(1)

	if len(params) == 3 {
		min = int16(params[1])
		max = int16(params[2])
	}

	rand.Seed(time.Now().UnixNano())
	value := min + int16(rand.Int()) * (max - min)

	time.Sleep(1 * time.Millisecond)

	return value
}

func Int8Generator(params ...int) int8 {
	min := int8(0)
	max := int8(1)

	if len(params) == 3 {
		min = int8(params[1])
		max = int8(params[2])
	}

	rand.Seed(time.Now().UnixNano())
	value := min + int8(rand.Int()) * (max - min)

	time.Sleep(1 * time.Millisecond)

	return value
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

// Others
func UintGenerator(params ...int) uint {
	min := uint(0)
	max := uint(1)

	if len(params) == 3 {
		min = uint(params[1])
		max = uint(params[2])
	}

	rand.Seed(time.Now().UnixNano())
	value := min + uint(rand.Uint64()) * (max - min)

	return value
}

func Uint64Generator(params ...int) uint64 {
	min := uint64(0)
	max := uint64(1)

	if len(params) == 3 {
		min = uint64(params[1])
		max = uint64(params[2])
	}

	rand.Seed(time.Now().UnixNano())
	value := min + uint64(rand.Uint64()) * (max - min)

	return value
}

func Uint32Generator(params ...int) uint32 {
	min := uint32(0)
	max := uint32(1)

	if len(params) == 3 {
		min = uint32(params[1])
		max = uint32(params[2])
	}

	rand.Seed(time.Now().UnixNano())
	value := min + uint32(rand.Uint32()) * (max - min)

	return value
}

func Uint16Generator(params ...int) uint16 {
	min := uint16(0)
	max := uint16(1)

	if len(params) == 3 {
		min = uint16(params[1])
		max = uint16(params[2])
	}

	rand.Seed(time.Now().UnixNano())
	value := min + uint16(rand.Uint32()) * (max - min)

	return value
}

func Uint8Generator(params ...int) uint8 {
	min := uint8(0)
	max := uint8(1)

	if len(params) == 3 {
		min = uint8(params[1])
		max = uint8(params[2])
	}

	rand.Seed(time.Now().UnixNano())
	value := min + uint8(rand.Uint32()) * (max - min)

	return value
}

func DumpToFile(file_name string, content string) {
	file, err := os.OpenFile(file_name, os.O_TRUNC | os.O_CREATE | os.O_WRONLY, 0644)

	if err != nil {
		panic("Error creating file")
	}

	defer file.Close()

	_, err = fmt.Fprintln(file, content)

	if err != nil {
		panic("Error writing to file")
	}
}	