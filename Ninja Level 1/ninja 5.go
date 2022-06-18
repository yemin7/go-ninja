package main

import "fmt"

var y int = 25

type onnut int

var x onnut

func main() {
	fmt.Printf("%v\t%T\n", x, x)
	x = 7
	fmt.Println("X value = ", x)

	y = int(x)
	fmt.Println("Coversion x value to y: ", y)
}
