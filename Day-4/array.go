package main

import "fmt"

func main() {
	num_array := []int{2, 1, 5, 4, 3, 7}
	fmt.Println(num_array)
	max := 0
	min := 0

	for i := 0; i < len(num_array); i++ {
		if num_array[i] > max {
			max = num_array[i]
		}
	}
	fmt.Println("Max value: ", max)

	for i := len(num_array) - 1; i > 0; i-- {
		if num_array[i] < num_array[i-1] {
			min = num_array[i]
		}
	}
	fmt.Println("Min value: ", min)
}
