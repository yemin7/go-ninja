package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "みな",
		Age:   26,
	}

	u2 := user{
		First: "さな",
		Age:   25,
	}

	u3 := user{
		First: "もも",
		Age:   26,
	}

	users := []user{u1, u2, u3}

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)
	fmt.Println(string(bs))

}
