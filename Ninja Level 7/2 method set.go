package main

import "fmt"

type person struct {
	first, last string
}

func (jp *person) jp_line() {
	fmt.Println("It is method", jp.first, jp.last)
	jp.first = "もも"
	jp.last = "平井"
	fmt.Println("After change in method: ", jp.first, jp.last)
}

type jp_singer interface {
	jp_line()
}

func main() {
	p1 := person{
		first: "みな",
		last:  "名井",
	}
	//	p1.jp_line()
	changeme(&p1)
	fmt.Printf("%T\n", &p1)
	fmt.Println("After changeme", p1)
}

func changeme(jp jp_singer) {
	jp.jp_line()
	//	fmt.Println("Before", p.first, p.last)
	//	p.first = "さな"
	//	p.last = "湊崎"
	//	fmt.Println("After", p.first, p.last)
}
