package main

import "fmt"

func main() {
	c := make(chan int, 1)
	//	c := make(chan int, 2)
	c <- 25
	//	c <- 7
	fmt.Println(<-c)
	//	fmt.Println(<-c)
}
