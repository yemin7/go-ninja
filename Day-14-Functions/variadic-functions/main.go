package main

import "fmt"

func main() {
	a1 := average(34, 57, 24, 18, 85, 37)
	fmt.Println(a1)
}

func average(a ...int) float64 {
	fmt.Printf("%T\n", a)
	var total int
	for _, v := range a {
		total += v
	}
	return float64(total / len(a))
}
