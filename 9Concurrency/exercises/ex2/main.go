package main

import (
	"fmt"
)

//create a type person struct
type person struct {
	first string
}

//create a type human interface
// to implicitly implement the interface, a human must have the speak method
type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "Jana",
	}
	saySomething(&p1)
	//saySomething(p1) => does not work
	//uses interface
}

//attach a method speak to type person using a pointer receiver
func (p *person) speak() {
	fmt.Println("Hello")
}

//create func “saySomething”
//have it take in a human as a parameter
// have it call the speak method
func saySomething(h human) {
	h.speak()
}

//show the following in your code
//○ you CAN pass a value of type *person into saySomething
//○ you CANNOT pass a value of type person into saySomething
func saySomethingP(p *person) {
	p.speak()
}

//https://play.golang.org/p/FAwcQbNtMG
// type circle struct {
// 	radius float64
// }

// type shape interface {
// 	area() float64
// }

// func (c *circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func info(s shape) {
// 	fmt.Println("area", s.area())
// }

// func main2() {
// 	c := circle{5}
// 	// info(c)
// 	info(&c)
// 	// fmt.Println(c.area())
// }
