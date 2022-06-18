package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bs)
	loginPass := `password1234`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPass))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("You're login.")
}
