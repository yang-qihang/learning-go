package main

import "fmt"

func main() {
	// Exercise 1
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	greetings1 := greetings[:2]
	greetings2 := greetings[1:4]
	greetings3 := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(greetings1)
	fmt.Println(greetings2)
	fmt.Println(greetings3)

	// Exercise 2
	message := "Hi 👩 and 👨"
	runes := []rune(message)
	// In fact you don't actually need to convert the rune to a string first.
	// runes := []rune(message)
	// fmt.Printf("%c\n", runes[3])
	fmt.Println(string(runes[3]))

	// Exercise 3
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	emp1 := Employee{"Person", "Personson", 12345}
	emp2 := Employee{
		firstName: "Mary",
		lastName:  "Jones",
		id:        65432,
	}
	var emp3 Employee
	emp3.firstName = "Pat"
	emp3.lastName = "Carter"
	emp3.id = 98765

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}
