package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	// The most common reason for iterating over the key is when a map is
	// being used as a set. In those situations, the value is unimportant.

	uniqueNames := map[string]bool{
		"Fred":  true,
		"Raul":  true,
		"Wilma": true,
	}

	for k := range uniqueNames {
		fmt.Println(k)
	}
}
