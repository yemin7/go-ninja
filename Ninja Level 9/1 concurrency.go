package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo: ", i)
		//time.Sleep(time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for j := 0; j < 5; j++ {
		fmt.Println("bar: ", j)
		//		time.Sleep(time.Millisecond)
	}
	wg.Done()
}
