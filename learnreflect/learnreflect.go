package learnreflect

import (
	"fmt"
	"reflect"
)

func Example1() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())
}

func Example2() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}

func Example3() {
	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("interface value:", v.Interface())
	fmt.Printf("interface value: %c\n", v.Interface())
	fmt.Println("kind is uint8:", v.Kind() == reflect.Uint8)
	x = uint8(v.Uint())
}
