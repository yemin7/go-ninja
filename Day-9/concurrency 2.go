package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//	fmt.Println("CPUs\t\t", runtime.NumCPU())
	//	fmt.Println("Goroutine\t", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()
	//	fmt.Println("CPUs\t\t", runtime.NumCPU())
	//	fmt.Println("Goroutine\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
