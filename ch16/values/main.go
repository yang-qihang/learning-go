package main

import (
	"fmt"
	"reflect"
)

func main() {
	// How to read value back out of a `reflect.Value`
	s := []string{"a", "b", "c"}
	sv := reflect.ValueOf(s)
	s2 := sv.Interface().([]string)
	fmt.Println(s2)

	// Use reflection to set the value of a variable
	i := 10
	iv := reflect.ValueOf(&i)
	ivv := iv.Elem()
	ivv.SetInt(20)
	fmt.Println(i) // prints 20

	fmt.Println(reflect.ValueOf(i).CanSet()) // prints false
}
