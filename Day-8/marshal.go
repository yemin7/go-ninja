package main

import (
	"encoding/json"
	"fmt"
)

type singer struct {
	First, Last string
}

func main() {
	s1 := singer{
		First: "みな",
		Last:  "名井",
	}
	s2 := singer{
		First: "さな",
		Last:  "湊崎",
	}

	jp_singer := []singer{s1, s2}
	fmt.Printf("%T\n", jp_singer)
	fmt.Println(jp_singer)

	bs, err := json.Marshal(jp_singer)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(bs))
}
