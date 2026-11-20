package main

import "fmt"

type Number interface {
	int | float32 | float64
}

func Double[T Number](val T) T {
	return val * 2
}

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type MyInt int

// Alternatively, you can choose to use:
// return strconv.Itoa(int(mi))
//
// Under the hood,
// calling fmt.Sprint on a type implementing Stringer
// can invoke that type's String() method.

func (mi MyInt) String() string {
	return fmt.Sprint(int(mi))
}

func PrintPrintable[T Printable](pt T) {
	fmt.Println(pt.String())
}

func main() {
	var myValue MyInt = 3
	PrintPrintable(myValue)
}
