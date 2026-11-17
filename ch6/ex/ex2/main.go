package main

import "fmt"

func UpdateSlice(sliceStr []string, str string) {
	sliceStr[len(sliceStr)-1] = str
	fmt.Println(sliceStr)
}

func GrowSlice(sliceStr []string, str string) {
	sliceStr = append(sliceStr, str)
	fmt.Println(sliceStr)
}

func main() {
	test := []string{"apple", "banana", "orange", "cake"}
	fmt.Println(test)

	UpdateSlice(test, "peach")
	fmt.Println(test)

	GrowSlice(test, "pig")
	fmt.Println(test)
}
