package main

import "fmt"

var y int

type mina int

var m mina

func main() {
	y = 25
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	m = 26
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	y = int(m)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
