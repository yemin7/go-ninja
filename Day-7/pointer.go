package main

import "fmt"

func main() {

	a := 25
	fmt.Println(a)
	fmt.Println(&a) // gives the address
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	//fmt.Println(&b)
	fmt.Println(*b) //give the value stored at the address
	fmt.Println(*&a)

	/*	var b int = a
		fmt.Println(b)
		fmt.Println(&b)

		var c *int = &a
		fmt.Println(c)
		fmt.Println(&c)

		d := &a
		fmt.Println(d)
		fmt.Println(&d)
		fmt.Println(*d)
	*/
}
