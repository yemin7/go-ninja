package main

import "fmt"

func main() {
	num := make([]int, 6, 12)
	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))

	num[5] = 35
	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))

	num = append(num, 36)
	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))

	appendx := []int{5, 7, 10, 13, 29, 30}
	num = append(num, appendx...)
	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))
}
