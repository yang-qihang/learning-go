package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(divAndRemainder(5, 2))
}

func divAndRemainder(num, denom int) (result int, remainder int, err error) {
	// assign some values
	result, remainder = 20, 30
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}
