package main

import "fmt"

func main() {
	//anonymous struct
	//use if struct is only used once
	//put struct directly in p1
	p1 := struct {
		first string
		last  string
	}{
		first: "James",
		last:  "Bond",
	}
	fmt.Println(p1)
}

//before: non anonymous because it has a name "person"
// type person struct {
// 	first string
// 	last  string
// }

// func main() {

// 	p1 := person{
// 		first: "James",
// 		last:  "Bond",
// 	}
// 	fmt.Println(p1)
// }
