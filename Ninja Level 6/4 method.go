package main

import "fmt"

type person struct {
	first, last string
	age         int
}

func (p person) speak() {
	fmt.Println("My name is ", p.first, p.last, "Age: ", p.age)
}

func main() {
	p1 := person{
		first: "Mina",
		last:  "Myoi",
		age:   26,
	}
	p1.speak()

}
