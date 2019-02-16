package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	info(c)
	//info(&c) works because receier t can get t or *t but pointers => value pointer
}

//NON-POINTER RECEIVER & POINTER VALUE
// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// } ==> non-pointer

// func info(s shape) {
// 	fmt.Println("area", s.area())
// }

// func main() {
// 	c := circle{5}
// 	info(&c)
// } ==>pointer

//POINTER RECEIVER & POINTER VALUE
// func (c *circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// } ==> pointer

// func info(s shape) {
// 	fmt.Println("area", s.area())
// }

// func main() {
// 	c := circle{5}
// 	info(&c) //==> pointer
// }
