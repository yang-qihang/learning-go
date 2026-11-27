package main

import (
	"fmt"
	"math"
	"sync"
)

func initResult() map[int]float64 {
	res := map[int]float64{}
	for i := range 100_000 {
		res[i] = math.Sqrt(float64(i))
	}
	return res
}

var initResultCached func() map[int]float64 = sync.OnceValue(initResult)

func lookUpSqrtOfMultiplesOf1000() {
	for i := range 100 {
		fmt.Println("Num:", i*1000, "square root:", initResultCached()[i*1000])
	}
}

func main() {
	lookUpSqrtOfMultiplesOf1000()
}
