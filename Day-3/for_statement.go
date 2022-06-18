package main

import "fmt"

func main() {

	//Like a C for
	//for init; condition; post {}
	for num1 := 1; num1 <= 5; num1++ {
		fmt.Println(num1)
	}
	fmt.Println("First loop is done.")
	fmt.Println("")

	//Like a C while
	//for condition {}
	num2 := 1
	for num2 <= 5 {
		fmt.Println(num2)
		num2++
	}
	fmt.Println("Second loop is done.")
	fmt.Println("")

	//Like a C for (;;)
	//for {}
	num3 := 1
	for {
		if num3 > 5 {
			break
		}
		fmt.Println(num3)
		num3++
	}
	fmt.Println("Third loop is done.")
}
