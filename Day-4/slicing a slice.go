package main

import "fmt"

func main() {
	num := []int{2, 5, 8, 1, 9}

	fmt.Println(num)
	fmt.Println(num[1:])
	fmt.Println(num[1:3])

	k := 2
	num = append(num[:k], num[(k+1):]...)
	fmt.Println(num)
}
