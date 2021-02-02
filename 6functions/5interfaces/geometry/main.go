package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

//Interfaces are named collections of method signatures.
type geometry interface {
	area() float64
	perim() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

//If a variable has an interface type, then we can call methods that are in the named interface.
//Here’s a generic measure function taking advantage of this to work on any geometry.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())

}

func main() {
	c1 := circle{5.0}
	r1 := rectangle{3.0, 5.0}

	measure(c1)
	measure(r1)

}

//The circle and rect struct types both implement the geometry interface
//so we can use instances of these structs as arguments to measure.
