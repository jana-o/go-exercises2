package main

import "fmt"

func main() {
	var x [5]int
	x[2] = 1
	fmt.Println(x)

	twoD()
}
func twoD() {
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

//multi dimensional array: size cannot vary, but with slices
//output:
// [[0 1 2 3 4] [1 2 3 4 5] [2 3 4 5 6]]

//twoD[i][j] = j
//[[0 1 2 3 4] [0 1 2 3 4] [0 1 2 3 4]]
