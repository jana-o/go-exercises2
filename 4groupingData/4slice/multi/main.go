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
}
