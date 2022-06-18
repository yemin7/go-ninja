package main

import (
	"fmt"
)

func main() {

	//using func literal
	/*
		c := make(chan int)
		go func() {

			c <- 42
		}()
	*/

	//using buffer
	c := make(chan int, 1)
	c <- 42

	fmt.Println(<-c)
}
