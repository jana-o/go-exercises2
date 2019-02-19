package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

// Create a custom error message using “fmt.Errorf”.
func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		fmt.Println(err)
		return
		//or fatal and then no return bec program ends
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		return []byte{}, fmt.Errorf("there was an error: %v", err)
		// return []byte{}, fmt.Println("there was an error: %v", err)
		//does not work bc Println returns n int and err => we only need err
		//fatal does not work bec no error returned but we need to return here

	}
	return bs, nil
}
