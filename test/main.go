// This lab aims to test if assigning a concrete value to an interface in Go
// involves a copy of the concrete value.
package main

import "fmt"

type SomeType struct {
	value int
}

func (s SomeType) getValue() int {
	return s.value
}

type SomeInterface interface {
	getValue() int
}

func main() {
	myValue := SomeType{value: 10}

	var myInterface SomeInterface = myValue

	myValue.value = 20

	fmt.Println(myValue.getValue())     // 20
	fmt.Println(myInterface.getValue()) // 10
}
