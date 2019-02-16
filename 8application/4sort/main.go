package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println("unsorted:", xi)
	sort.Ints(xi)

	//sort ints
	// https://godoc.org/sort#Ints
	//we do not need to assign it so anything bc nothing returned
	fmt.Println("sorted:", xi)
	fmt.Println("------")

	//sort strings
	fmt.Println("unsorted:", xs)
	sort.Strings(xs)
	fmt.Println("sorted", xs)
}
