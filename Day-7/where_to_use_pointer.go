package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x value: ", x)
	fmt.Println("x address: ", &x)
	foo(&x)
	fmt.Println("x value after foo: ", x)
}

func foo(y *int) {
	fmt.Println(y)
	fmt.Println("Y value: ", *y)
	fmt.Println("Y address: ", &y)
	*y = 26
	fmt.Println("Y value after: ", *y)
	fmt.Println("Y address after: ", &y)
}
