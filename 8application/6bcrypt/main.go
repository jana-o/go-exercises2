package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	//https://godoc.org/golang.org/x/crypto/bcrypt
	// GenerateFromPassword(password []byte, cost int) ([]byte, error)

	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	//func CompareHashAndPassword(hashedPassword, password []byte) error
	//hashedPassword = bs
	loginPword1 := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("you can't login")
		return
	}
	fmt.Println("You are logged in")
}
