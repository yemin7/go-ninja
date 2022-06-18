package main

import "fmt"

func main() {
	ninja := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(ninja))
	fmt.Println(even(ninja)())

	fmt.Println(odd(sum, ninja...))
}

func sum(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(y []int) func() int {
	var total []int
	for _, v := range y {
		if v%2 == 0 {
			total = append(total, v)
		}
	}
	fmt.Println("Even numbers: ", total)
	return func() int {
		sum := 0
		for _, v := range total {
			sum += v
		}
		return sum
	}
}

func odd(f func(x []int) int, y ...int) int {
	var total []int
	for _, v := range y {
		if v%2 != 0 {
			total = append(total, v)
		}
	}
	fmt.Println("Odd number: ", total)
	return f(total)
}
