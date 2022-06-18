package main

import "fmt"

const (
	_      = iota
	year_1 = iota * 365
	year_2 = iota * 365
	year_3 = iota * 365
	year_4 = iota * 365
)

func main() {
	fmt.Println("Days in First Year:\t", year_1)
	fmt.Println("Days in Second Year:\t", year_2)
	fmt.Println("Days in Third Year:\t", year_3)
	fmt.Println("Days in Fourth Year:\t", year_4)

}
