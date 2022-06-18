package main

import "fmt"

func main() {

	fmt.Println("Hello main.")
	foo()

	bar("Hello from bar.")

	str_return := returnfunc("Return Function.")
	fmt.Println(str_return)

	return1, return2 := multireturn("Mina", "Dahyun")
	fmt.Println(return1, return2)
}

func foo() {
	fmt.Println("Hello from foo.")
}

//Function Pass by Value
func bar(str1 string) {
	fmt.Println(str1)
}

//Return Function
func returnfunc(str1 string) string {
	return fmt.Sprint("Hello from ", str1)
}

func multireturn(first, last string) (string, bool) {
	first_return := fmt.Sprint(first, " and ", last, ` said "Hello Once!"`)
	second_return := true
	return first_return, second_return
}
