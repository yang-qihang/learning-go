package main

import (
	"errors"
	"fmt"
)

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

func main() {
	result, err := doubleEven(1)
	if err != nil {
		// Here Println automatically calls the Error method on the error value
		// automatically.
		fmt.Println(err)
	}
	fmt.Println(result)
}
