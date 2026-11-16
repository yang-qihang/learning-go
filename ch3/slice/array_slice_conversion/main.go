package main

import "fmt"

func main() {
	// Converting Arrays to Slices
	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]

	x[0] = 10
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	fmt.Println()

	// Converting Slices to Arrays
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 10

	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)
}
