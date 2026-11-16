package main

import "fmt"

func main() {
	var total int

	// This can be modernized to be for i := range 10
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}

	// This exercise is intended to show that unintentionally use =
	// where a := is expected can shadow the variables and result in
	// some surprising errors
	fmt.Println(total)
}
