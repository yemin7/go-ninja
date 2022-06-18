package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int

var wg sync.WaitGroup

const gs = 10

func main() {
	wg.Add(1)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("Total counter:", counter)
}

func foo() {
	for i := 0; i < gs; i++ {
		v := counter
		runtime.Gosched()
		v++
		counter = v
		fmt.Println("foo: ", v, "counter: ", counter)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < gs; i++ {
		v := counter
		runtime.Gosched()
		v++
		counter = v
		fmt.Println("bar: ", v, "counter: ", counter)
	}
	wg.Done()
}
