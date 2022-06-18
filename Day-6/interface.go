package main

import "fmt"

type person struct {
	first, last string
	age         int
}

type jp_line struct {
	person
	japan bool
}

//method
func (jp jp_line) speak() {
	fmt.Println(jp.first, " ", jp.last, ` said "Hello Once!"`)
}

//func (p person) speak() {
//	fmt.Println("I am ", p.first, p.last)
//}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I am human", h.(jp_line).first)
}

func main() {

	singer1 := jp_line{
		person: person{
			first: "もも",
			last:  "平井",
			age:   26,
		},
		japan: true,
	}
	fmt.Println(singer1)
	singer1.speak() //method

	bar(singer1)
}
