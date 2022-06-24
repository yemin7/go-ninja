package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "inner scope"
		fmt.Println(y)
	}
	//	fmt.Println(y)	//outside scope of y
}
