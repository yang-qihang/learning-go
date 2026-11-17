package main

import "fmt"

func prefixer(prefix string) func(string) string {
	return func(s string) string {
		return prefix + " " + s

		// The follow is OK, but can be simplied considerably...
		// prefixSlice := append([]rune(prefix), []rune(" ")...)
		// completeSlice := append(prefixSlice, []rune(s)...)
		// return string(completeSlice)
	}
}

func main() {
	helloPrefix := prefixer("hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
