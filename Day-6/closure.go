package main

import "fmt"

func main() {
	x := 1
	{
		y := 25
		fmt.Println(y)
	}
	fmt.Println(x)
}
