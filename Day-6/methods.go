package main

import "fmt"

type person struct {
	first, last string
	age         int
}

type japanese struct {
	person
	jp bool
}

func (jp japanese) speak() {
	fmt.Println(jp.first, " ", jp.last, ` said "Hello Once!"`)
}

func main() {
	jp1 := japanese{
		person: person{
			first: "もも",
			last:  "平井",
			age:   26,
		},
		jp: true,
	}
	fmt.Println(jp1)
	jp1.speak()
}
