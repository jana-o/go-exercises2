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
	return c.radius * c.radius * math.Pi
}

func info(sp shape) {
	fmt.Println(sp.area())
}

func main() {
	c := circle{2.0}
	info(c) //value receiver, value type
}
