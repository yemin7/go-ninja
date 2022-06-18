package main

import "fmt"

func main() {
	f := func(x int) {
		fmt.Println("Function expression.", x)
	}
	f(25)
}
