package main

import "fmt"

func main() {
	num := 25
	fmt.Printf("%d\t%b\t%#x\n", num, num, num)

	num_shift := num << 1
	fmt.Printf("%d\t%b\t%#x\n", num_shift, num_shift, num_shift)
}
