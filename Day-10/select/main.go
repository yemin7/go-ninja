package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	//send
	go send(even, odd, quit)

	//receive
	receive(even, odd, quit)
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
	//	close(e)
	//	close(o)
	//	q <- 0
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from even channel:", v)
		case v := <-o:
			fmt.Println("from odd channel:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from comma ok ", i, ok)
			} else {
				fmt.Println("from comma ok quit", i)
			}
			return
		}
	}
}

//func receive(e, o, q <-chan int) {
//	for {
//		select {
//		case v := <-e:
//			fmt.Println("from even channel:", v)
//		case v := <-o:
//			fmt.Println("from odd channel:", v)
//		case i, ok := <-q:
//			if !ok {
//				fmt.Println("from comma ok", ok)
//			} else {
//				fmt.Println("from quit channel:", i)
//			}
//			return
//		}
//	}
//}
