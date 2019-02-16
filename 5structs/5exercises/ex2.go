package main

import "fmt"

type vehicle struct {
	door  int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			door:  2,
			color: "red",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			door:  4,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println(t)
	fmt.Println(t.door)

	fmt.Println(s)
	fmt.Println(s.color)

	x := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
		},
		favDrinks: []string{
			"Martini",
			"Vodka",
		},
	}
	fmt.Println(x.friends)

	for k, v := range x.friends {
		fmt.Println(k, v)
	}
}
