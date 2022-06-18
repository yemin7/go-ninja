package main

import "fmt"

func main() {
	ninja := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	ninja = append(ninja[:3], ninja[6:]...)
	fmt.Println(ninja)
}
