package main

import "fmt"

func main() {
	num := []int{2, 5, 8, 1, 9}
	fmt.Println(num)
	max := 0

	//i = index, v= value
	for _, v := range num {
		//		fmt.Printf("Index: %v, Value: %v\n", i, v)

		if v > max {
			max = v
		}
	}
	fmt.Println("Max value:", max)

}
