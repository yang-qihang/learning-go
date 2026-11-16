package main

import "fmt"

func main() {
	// map read & write
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2

	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kitten"])

	totalWins["Kitten"]++
	fmt.Println(totalWins["Kitten"])

	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])

	// the comma ok idiom
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	delete(m, "hello")

	// Using Maps as Sets
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[100])

	if intSet[100] {
		fmt.Println("100 is in the set")
	}
	// Some people prefer to use struct{} when a map is being used to implement a set.
}
