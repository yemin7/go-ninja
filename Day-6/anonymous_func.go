package main

import "fmt"

func main() {
	foo()
	//anonymous Function
	func() {
		fmt.Println("Hello anonymous.")
	}()

	x := 42
	func(y int) {
		fmt.Println("X value: ", y)
	}(x)

}

func foo() {
	fmt.Println("Hello from foo.")
}
