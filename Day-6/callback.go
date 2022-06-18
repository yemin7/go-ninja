package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(i...))
	fmt.Println(even(i...))

	fmt.Println(odd(sum, i...))
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(x ...int) []int {
	even := []int{}
	for _, v := range x {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	return even
}

func odd(f func(x ...int) int, y ...int) int {
	var odd []int
	for _, v := range y {
		if v%2 != 0 {
			odd = append(odd, v)
		}
	}
	//	fmt.Println(odd)
	return f(odd...)
}
