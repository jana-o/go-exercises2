package main

import (
	"fmt"
	"math"
)

//create a type SQUARE and type CIRCLE
type circle struct {
	radius float64
}

type square struct {
	length float64
}

//attach method to each that calculates AREA and returns it
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
	//return math.Pow(s.length, 2)
}

//create a type SHAPE that defines an interface as anything that has the AREA method
type shape interface {
	area() float64
} //need float64 here also because area()

//create a func INFO which takes type shape and prints the area
func info(s shape) {
	fmt.Println(s.area())
}

//create a value of type square and circle
//use func infor to print area of sqaure and circle
func main() {
	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
	}

	info(circ)
	info(squa)
}
