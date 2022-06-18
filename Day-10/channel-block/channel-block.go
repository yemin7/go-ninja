package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	go send(channel1)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
}

func send(c chan string) {

	reportNap("sending goroutine", 2)
	fmt.Println("*** sending value ***")
	c <- "a"
	fmt.Println("*** sending value ***")
	c <- "b"

}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wake up")
}
