package main

import "fmt"

func main() {
	c := make(chan int)
	go foo(c)
	bar(c)
}

func foo(a chan<- int) {
	a <- 25
}
func bar(b <-chan int) {
	fmt.Println(<-b)
}
