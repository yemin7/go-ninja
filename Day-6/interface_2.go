package main

import "fmt"

type singer struct {
	first, last string
	age         int
}

type japan struct {
	singer
	japanese bool
}

func (jp japan) speak() {
	fmt.Println(jp.first, jp.last, `said "Hello Once!"`)
}

type human interface {
	speak()
}

func dance(h human) {
	fmt.Println(h.(japan).first, "Dance!!!")
}

func main() {

	singer1 := japan{
		singer: singer{
			first: "もも",
			last:  "平井",
			age:   26,
		},
		japanese: true,
	}

	fmt.Println(singer1)
	singer1.speak()
	dance(singer1)
}
