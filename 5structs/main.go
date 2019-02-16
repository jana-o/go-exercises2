//structure, aggregate data type
//similar to object/class
//allows you to aggregate values of different types

package main

import "fmt"

// person{} => composite literal
type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	fmt.Println(sa1)

	p2 := person{
		first: "Miss",
		last:  "Moneypenn",
		age:   27,
	}
	fmt.Println(p2)
	//inner type getx promoted to the outer type
	// no need to write sa1.person.first  => sa1.first is ok
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)

}
