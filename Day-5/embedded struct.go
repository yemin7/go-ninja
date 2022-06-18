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

func main() {

	jp1 := japanese{
		person: person{
			first: "もも",
			last:  "平井",
			age:   26,
		},
		jp: true,
	}

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

	fmt.Printf("Japanese Singer: %v %v\nage:\t%v\n", jp1.first, jp1.last, jp1.age)
}
