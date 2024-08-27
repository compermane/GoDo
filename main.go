package main

import (
	"fmt"
	"reflect"
)

// Function é um tipo que pode armazenar qualquer função
type Function interface{}

// ApplyOperation aplica uma função genérica com argumentos variáveis
func ApplyOperation(f Function, args ...interface{}) []interface{} {
    fn := reflect.ValueOf(f)
    if len(args) != fn.Type().NumIn() {
        panic("Número incorreto de argumentos fornecidos")
    }

    in := make([]reflect.Value, len(args))
    for i, arg := range args {
        in[i] = reflect.ValueOf(arg)
    }

    results := fn.Call(in)
    out := make([]interface{}, len(results))
    for i, result := range results {
        out[i] = result.Interface()
    }

    return out
}

// Algumas funções para teste
func add(x, y int) int {
    return x + y
}

func concat(a, b string) string {
    return a + b
}

func main() {
    // Aplicar operação add
    sum := ApplyOperation(add, 3, 4)
    fmt.Printf("Sum: %v\n", sum[0]) // Saída: Sum: 7

    // Aplicar operação concat
    combined := ApplyOperation(concat, "Hello, ", "World!")
    fmt.Printf("Concatenated: %v\n", combined[0]) // Saída: Concatenated: Hello, World!
}
