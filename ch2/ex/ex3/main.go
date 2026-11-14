package main

import (
	"fmt"
	"math"
)

func main() {
	var b byte = math.MaxUint8
	var i int32 = math.MaxInt32
	var u uint64 = math.MaxUint64

	// It can also be calculated manually
	// var b byte = (1 << 8) - 1
	// var i int32 = (1 << 31) - 1
	// var u uint64 = (1 << 64) - 1

	fmt.Println(b)
	fmt.Println(i)
	fmt.Println(u)
}
