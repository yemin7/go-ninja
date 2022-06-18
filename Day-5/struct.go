package main

import "fmt"

type person struct {
	first, last string
	age         int
}

func main() {
	person1 := person{
		first: "みな",
		last:  "名井",
		age:   26,
	}
	person2 := person{
		first: "さな",
		last:  "湊崎",
		age:   25,
	}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.first, person1.last, person1.age)
	fmt.Println(person2.first, person2.last, person2.age)
}
