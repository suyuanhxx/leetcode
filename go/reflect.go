package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.14
	fmt.Println("Type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("Type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Println(v.Interface().(float64))
	fmt.Println(v.CanSet())

	v = reflect.ValueOf(&x)
	fmt.Println(v.CanSet())
	v = v.Elem()
	fmt.Println(v.CanSet())
	v.SetFloat(3.141592)
	fmt.Println(v.Interface())
	fmt.Println(x)

}
