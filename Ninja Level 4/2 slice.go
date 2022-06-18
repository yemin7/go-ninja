package main

import "fmt"

func main() {
	ninja := []int{3, 46, 28, 94, 9, 83, 78, 34, 59, 92}
	for k, v := range ninja {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
	fmt.Printf("%T\n", ninja)
}
