package main

import "fmt"

type Struct1 struct {
	name  string
	value int
}

type Struct2 Struct1
type Struct3 struct {
	field1 Struct1
	field2 Struct2
}

func main() {
	var s1 = Struct1{}
	var s2 = Struct2{}
	s1 = struct {
		name  string
		value int
	}{}
	s2 = struct {
		name  string
		value int
	}{}
	s3 := struct {
		field1 Struct1
		field2 Struct2
	}{}
	fmt.Println(s1, s2, s3)
}
