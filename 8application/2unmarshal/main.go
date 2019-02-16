package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

//https://mholt.github.io/json-to-go/
//unmarshal: if have a JSON and i want it to turn into a Go data structure

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all the data", people)

	for i, v := range people {
		fmt.Println("\nPerson Number", i)
		fmt.Println(v.First, v.Last, v.Age)

	}
}
