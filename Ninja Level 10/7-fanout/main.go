package main

import "fmt"

func main() {
	channel1 := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				channel1 <- j
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-channel1)
	}
}
