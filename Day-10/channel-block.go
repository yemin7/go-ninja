package main

import "fmt"

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go abc(channel1)
	abc_r(channel1)
	go def(channel2)
}

func abc_r(c <-chan string) {
	for v := range c {
		fmt.Println(v)
	}
}

func abc(c chan<- string) {
	c <- "a"
	c <- "b"
	c <- "c"
	close(c)
}

func def(c chan string) {
	c <- "d"
	c <- "e"
	c <- "f"
}
