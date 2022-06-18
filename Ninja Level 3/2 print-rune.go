package main

import "fmt"

func main() {
	x := 1
	for alpha := 65; alpha <= 90; alpha++ {
		fmt.Println(x)
		for num := 1; num <= 3; num++ {
			fmt.Printf("\t%U %q\n", alpha, alpha)
		}
		x++
	}
}
