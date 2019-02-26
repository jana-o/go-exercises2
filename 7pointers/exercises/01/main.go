package main

import "fmt"

func main() {
	age := 44
	fmt.Println("start", &age) //address

	changeMe(&age)

	fmt.Println("after", &age) //address
	fmt.Println("after", age)  //24

}

func changeMe(z *int) {
	fmt.Println("z before:", z)  // address
	fmt.Println("z before:", *z) //44

	*z = 24
	fmt.Println("z after:", z)   // address
	fmt.Println("*z after:", *z) //24
}

//without pointer:  changeMe(age)
// func changeMe(z int) {
// 	fmt.Println("z before:", &z) // second address
// 	fmt.Println("z before:", z)  //44

// 	z = 24
// 	fmt.Println("z after:", &z) // second address
// 	fmt.Println("*z after:", z) //24
// }
// in main()
// fmt.Println("after", &age) //first address
// 	fmt.Println("after", age)  //44
