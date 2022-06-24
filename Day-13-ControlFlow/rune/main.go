package main

import "fmt"

func main() {
	s := "Hello"
	fmt.Println([]byte(s))
	for _, v := range s {
		fmt.Println(string(v))
	}
}
