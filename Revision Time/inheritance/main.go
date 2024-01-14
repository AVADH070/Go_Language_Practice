package main

import "fmt"

type simple struct {
	name   string
	number int
}

type Inheritance struct {
	name   string
	number int
	simple
}

type allFunction interface {
	getName() string
}

func (i Inheritance) getName() string {
	return i.name
}

func (i simple) getName() string {
	return i.name
}

func calculate(a allFunction) {
	_, ok := a.(Inheritance)

	if ok {
		fmt.Print(a.getName())
	} else {
		fmt.Print(a.getName())
	}
}

func main() { 
	objI := Inheritance{
		name:   "Inheritance",
		number: 1234,
		simple: simple{
			name:   "Avadh Sonagara",
			number: 234,
		},
	}
	calculate(objI)
}
