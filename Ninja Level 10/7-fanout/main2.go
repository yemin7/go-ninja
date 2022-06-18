package main

import "fmt"

func main() {
	var x, y int = 10, 10
	c := fanout(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c)
	}
}

func fanout(x, y int) <-chan int {
	channel1 := make(chan int)
	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < y; j++ {
				channel1 <- j
			}
		}()
	}
	return channel1
}
