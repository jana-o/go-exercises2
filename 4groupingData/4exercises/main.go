package main

import "fmt"

func main() {
	fmt.Println("///ex 1: Arrays///")
	e1array()
	fmt.Println("///ex 2: Slices///")
	e2slices()
	fmt.Println("///ex 3: Slicing///")
	e3slicing()
	fmt.Println("///ex 4: append slice///")
	e4appendslices()
	fmt.Println("///ex 5: delete slice///")
	e5delete()
	fmt.Println("///ex 6: capacity, len slice///")
	e6states()
	fmt.Println("///ex 7: slice of slice///")
	e7sliceOfString()
	fmt.Println("///ex 8: map///")
	e8map()
}

func e1array() {
	// arr := [5]int{}
	// arr[0] = 1
	// arr[1] = 2

	// for i := range arr {
	// 	fmt.Println(arr[i])
	// }
	// fmt.Printf("%T\n%v", arr, arr)

	x := [5]int{42, 43, 44, 45, 46}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

func e2slices() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

func e3slicing() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x1 := append(x[:5])
	x2 := append(x[5:])
	x3 := append(x[2:7])
	x4 := append(x[1:6])
	fmt.Println(x)
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
}

func e4appendslices() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	//append a slice
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

}

func e5delete() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

func e6states() {
	// states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	x := make([]string, 50, 500)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// put a value into each of the 50 positions in the length of the slice
	// we are putting values into the 50 positions that are the length of the slice
	for i := 0; i < 50; i++ {
		x[i] = fmt.Sprintf("Position %d", i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// append values to the slice grows the length of the slice
	// the underlying array "capacity" of 500 is the same
	x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

func e7sliceOfString() {
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(x1)
	fmt.Println(x2)

	xx := [][]string{x1, x2}
	fmt.Println(xx)

	for i, x := range xx {
		fmt.Println("record:", i)
		for j, val := range x {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}

func e8map() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken`, `Martini`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	for i, v := range m {
		fmt.Printf("%v 's favorite things: \n", i)
		for j, val := range v {
			fmt.Printf("\t %v %v \n", j+1, val)

		}
	}

	//ex9: add a new record
	m[`ian`] = []string{`plants`, `books`}
	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}

	//ex10: delete a record
	delete(m, "no_dr")

	for i, v := range m {
		fmt.Println(i, v)
	}

}
