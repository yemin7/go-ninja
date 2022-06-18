package main

import "fmt"

//var c int

func main() {
	c := make(chan int)

	go foo(c)
	bar(c)
}

func foo(c chan<- int) {
	c <- 25
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}

//func foo() {
//	c = 2
//}
//
//func bar() {
//	fmt.Println(c)
//}
