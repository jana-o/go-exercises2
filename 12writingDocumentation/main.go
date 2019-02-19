package main

import (
	"exercises2/12writingDocumentation/dog"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}

	fmt.Println(fido)
}

//go run main.go
//godoc -http=:8080 check for dog package
