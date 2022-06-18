package main

import "fmt"

var ninja [5]int

func main() {
	ninja[0] = 3
	ninja[1] = 35
	ninja[2] = 64
	ninja[3] = 93
	ninja[4] = 73

	for k, v := range ninja {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}

	fmt.Printf("%T\n", ninja)
}
