// The observations made from writing this program can be also found
// in the Github repo for this book (the reference answer posted by the author).
//
// https://github.com/learning-go-book-2e/ch06/tree/main/exercise_solutions/ex3

package main

func main() {
	var people []Person
	//people := make([]Person, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		people = append(people, MakePerson("Fred", "Williamson", 25))
	}
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}
