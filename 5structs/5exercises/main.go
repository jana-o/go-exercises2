package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	icecreme  []string
}

func main2() {
	p1 := person{
		firstname: "James",
		lastname:  "Bond",
		icecreme: []string{
			"chocolate",
			"martini",
		},
	}
	p2 := person{
		firstname: "Miss",
		lastname:  "Moneypenny",
		icecreme: []string{
			"pistachio",
			"vanilla",
		},
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Printf("favorite icecreme: %v & %v", p1.icecreme, p2.icecreme)

	for i, v := range p1.icecreme {
		fmt.Println(i, v)
	}
	//ex2 store values of person in a map with key of lastname
	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}
	// fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.firstname)
		fmt.Println(v.lastname)

		for i, val := range v.icecreme {
			fmt.Println(i, val)
		}
		fmt.Println("----")
	}
}
