package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous")
	}()

	x := 25
	func(y int) {
		fmt.Println("y value: ", y)
	}(x)
}
