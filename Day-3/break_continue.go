package main

import "fmt"

func main() {
	num := 0
	for {
		num++
		if num > 10 {
			break
		}

		if num%2 != 0 {
			continue
		}
		fmt.Println(num)
	}
}
