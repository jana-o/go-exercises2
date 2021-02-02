package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini"}
	mp := []string{"Miss", "Moneypenny", "Strawberry"}
	fmt.Println(jb)
	fmt.Println(mp)

	//like an excel sheet
	xp := [][]string{jb, mp}
	fmt.Println(xp)

	multi()
}

func multi() {
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

//[[0] [1 2] [2 3 4]]
