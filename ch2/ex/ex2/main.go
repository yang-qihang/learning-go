package main

import "fmt"

func main() {
	const value = 1314
	// Here it relies on a rule called default typing for untyped constants.
	i := value
	var f float64 = value

	fmt.Printf("%d\n", i)
	fmt.Printf("%f\n", f)
}
