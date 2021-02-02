package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func totalArea(sp ...Shape) float64 {
	var total float64
	for _, s := range sp {
		total += s.area()
	}
	return total
}

func main() {
	s := Square{2.0}
	c := Circle{2.0}

	fmt.Println(totalArea(s, c))
	// fmt.Println(c.area())
	// fmt.Println(s.area())

}
