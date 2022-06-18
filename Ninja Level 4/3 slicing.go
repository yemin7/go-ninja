package main

import "fmt"

func main() {
	ninja := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(ninja[:5])
	fmt.Println(ninja[5:])
	fmt.Println(ninja[2:7])
	fmt.Println(ninja[1:6])

}
