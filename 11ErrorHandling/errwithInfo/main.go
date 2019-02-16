package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		//fatal will shut down program
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//when we call from pkg erros New => return &errorString() a pointer to a struct
		//errorString struct hast method and implements error interface
		//implicit error interface => polymorphism
		//thats why values can be of more than 1 type
		return 0, errors.New("norgate math: square root of negative number")
	}
	return 42, nil
}
