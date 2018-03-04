package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "  - the seretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "  - the person speak")
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

type human interface {
	speak() // hey baby if you got this method then you are my type -> person & secretagent
}

type fakeint int

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		"Mr.",
		"Robot",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)

	fmt.Println()
	//conversion
	var x fakeint = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
