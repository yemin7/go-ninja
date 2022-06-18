package main

import "fmt"

var y = 7

func main() {
	x := 25
	fmt.Printf("%b\n", x)  //base 2
	fmt.Printf("%o\n", x)  //base 8
	fmt.Printf("%x\n", x)  //base 16
	fmt.Printf("%U\n", x)  //unicode
	fmt.Printf("%#x\n", x) //unicode

	fmt.Println(y)
	fmt.Printf("%x\t%b\n", y, y)
	s := fmt.Sprintf("%x\t%b\n", y, y)
	fmt.Println(s)
}
