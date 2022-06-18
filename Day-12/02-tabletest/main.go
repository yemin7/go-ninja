package main

import "fmt"

func main() {
	fmt.Println("25 + 7 =", sumf(25, 7))

}

func sumf(s ...int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum + 1
}
