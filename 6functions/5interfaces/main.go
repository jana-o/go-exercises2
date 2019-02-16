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

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

//interfaces allow to define behaviour and allow polymorphism
//keyword (type) identifier (human) type (interface)
//a value can be of more than one type
//e.g. sa1 : is of type secretAgent but since it used speak() it's also type human
type human interface {
	speak()
}

// func bar(h human) {
// 	fmt.Println("I was passed into bar", h)
// }

//assertion
func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)

}

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
		first: "Dr.",
		last:  "No",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	//bar() can take in many types => polymorphism
	fmt.Println(p1)
	bar(sa1)
	bar(sa2)
	bar(p1)

}
