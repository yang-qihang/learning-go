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
		switch {
		case v%2 == 0:
			fmt.Println("Two")
		case v%3 == 0:
			fmt.Println("Three")
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		default:
			fmt.Println("Never mind")
		}
	}
}
