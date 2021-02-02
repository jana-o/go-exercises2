package main

import "fmt"

func main() {
	// x:= type{values} composite literal

	x := []int{4, 5, 6, 8, 42}
	fmt.Println(x)
	fmt.Println(x[2])
	fmt.Println(len(x))
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < 4; i++ {
		fmt.Println(i, x[i])
	} //both are same loops here

	///APPEND
	x = append(x, 77, 88)
	fmt.Println("append x: ", x)

	y := []int{123, 234, 56}
	//variadic ...  => unlimited number of this
	x = append(x, y...)
	fmt.Println("y: ", y)
	fmt.Println("appended y to x:", x)

	//DELETE from a slice: leave out 3,4
	x = append(x[:2], x[4:]...)
	fmt.Println("delete", x)

	xi := []int{4, 5, 7, 8, 9, 42}
	for i, v := range xi {
		fmt.Println(i, v)
	}

}

//a SLICE allows you to group together values of the same TYPE.
