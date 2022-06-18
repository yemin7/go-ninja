package main

import "fmt"

const first_ascii int = 32

func main() {

	for num := first_ascii; num <= 122; num++ {

		fmt.Printf("%c\t", num)
	}
}
