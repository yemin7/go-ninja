package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Begin CPU: ", runtime.NumCPU())
	fmt.Println("Begin Goroutine: ", runtime.NumGoroutine())
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("foo:1st goroutine")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("bar:2nd goroutine")
	}()
	fmt.Println("Mid CPU: ", runtime.NumCPU())
	fmt.Println("Mid Goroutine: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("End CPU: ", runtime.NumCPU())
	fmt.Println("End Goroutine: ", runtime.NumGoroutine())
}
