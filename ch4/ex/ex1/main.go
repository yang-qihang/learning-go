package main

import (
	"fmt"
	"math/rand"
)

func main() {

	randomIntegers := []int{}

	for i := 1; i <= 100; i++ {
		randomIntegers = append(randomIntegers, rand.Intn(100))
	}

	for _, v := range randomIntegers {
		fmt.Print(v, " ")
	}
}
