package main

import "fmt"

func main() {
	year := 1996
	for year < 2022 {
		fmt.Print(year, "\t")
		year++
	}
}
