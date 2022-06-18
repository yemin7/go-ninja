package main

import "fmt"

func main() {
	i := 0
	defer fmt.Println("defer", i)
	i++
	fmt.Println("i:", i)

	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}

	var x int
	x++
	fmt.Println("x:", x)
	j := c()
	fmt.Println("j:", j)
}

func c() (i int) {
	defer func() {
		i++
	}()
	return 1
}
