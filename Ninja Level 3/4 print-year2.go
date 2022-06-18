package main

import "fmt"

func main() {
	year := 1996
	for {
		if year == 2022 {
			break
		}
		fmt.Print(year, "\t")
		year++
	}
}
