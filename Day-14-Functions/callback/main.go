package main

import (
	"fmt"
)

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)
}

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, v := range numbers {
		if callback(v) {
			xs = append(xs, v)
		}
	}
	fmt.Printf("%T\n", callback)
	return xs
}
