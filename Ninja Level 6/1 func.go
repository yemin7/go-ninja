package main

import "fmt"

func main() {
	fmt.Println("Foo return:", foo())

	fmt.Println(bar())

}

func foo() int {
	return 25
}

func bar() (int, string) {
	return 2022, "This Year"
}
