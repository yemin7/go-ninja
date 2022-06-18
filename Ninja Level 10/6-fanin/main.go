package main

import "fmt"

func main() {
	channel1 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			channel1 <- i
		}
		close(channel1)
	}()

	for v := range channel1 {
		fmt.Println(v)
	}

}
