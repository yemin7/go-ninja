package main

import "fmt"

type onnut int

var x onnut

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 25
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
