package main

import (
	"encoding/json"
	"fmt"
)

// capital F bec needs to be exported

// type user struct {
// 	First string
// 	Age   int
// }

type user struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// marshal the []user to JSON
	//remember to ask yourself what you need to do to EXPORT a value from a package.

	//needs a pointer bec changes stuff
	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	err = json.Unmarshal([]byte(s), &users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
	//users and  not s! s shows just []

	//i, user not i, v
	for i, user := range users {
		fmt.Println("Person #", i)
		fmt.Println("\t", user.First, user.Last, user.Age)
		for _, saying := range user.Sayings {
			fmt.Println("\t", saying)
		}
	}

}
