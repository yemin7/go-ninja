package main

import "fmt"

func main() {
	num := []int{2, 5, 8, 1, 9}
	fmt.Println(num)
	num = append(num, 10, 15, 99)
	fmt.Println(num)

	num2 := []int{3, 6, 7, 1, 90}
	num = append(num, num2...) // use ... to append a slice to a slice
	fmt.Println("append slice", num)

	num = append(num[2:], num[:4]...)
	fmt.Println(num)
}
