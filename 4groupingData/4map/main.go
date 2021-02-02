package main

import "fmt"

func main() {
	// create an empty map: make(map[key-type]val-type)
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("This is the print", v)
	}

	//add element
	m["Todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	//slice of int
	xi := []int{4, 5, 7, 8, 9, 42}
	for i, v := range xi {
		fmt.Println(i, v)
	}

	//delete an item
	delete(m, "James")
	fmt.Println("delte James", m)
	//delete sth that does not exist works => check ok
	delete(m, "Ian Fleming")
	fmt.Println("delete Ian", m)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value", v)
		delete(m, "Miss Moneypenny")
	}

	fmt.Println(m)

	findElement(moons)
}

var moons = map[string]int{"Mercury": 0, "Venus": 0, "Earth": 1, "Mars": 2, "Jupiter": 67}

func findElement(mo map[string]int) {

	n, found := moons["Earth"]
	fmt.Println("Earth:", n, "Found:", found)

	n, found = moons["Neptune"]
	fmt.Println("Neptune:", n, "Found:", found)
}
