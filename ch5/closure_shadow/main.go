package main

import "fmt"

func main() {
	a := 20
	f := func() {
		fmt.Println(a)
		// this line is the key
		// If it has been a := 30 instead,
		// then it's more like shadowing!
		a = 30
		fmt.Println(a)
	}
	f()
	fmt.Println(a)
}
