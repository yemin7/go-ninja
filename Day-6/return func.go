package main

import "fmt"

func main() {
	func() {
		fmt.Println("anonymous function")
	}()
	fmt.Println(foo())
	//	b1 := bar()

	//	i := b1()
	//	fmt.Println(b1())
	fmt.Println(bar()())
}

func foo() string {
	return "Return foo"
}

func bar() func() int {
	return func() int {
		return 25
	}
}
