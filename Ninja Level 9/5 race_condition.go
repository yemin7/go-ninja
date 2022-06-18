package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

var count int

func main() {

	wg.Add(2)
	go race_condition("foo")
	go race_condition("bar")
	wg.Wait()
	fmt.Println("Total count: ", count)
}

func race_condition(s string) {
	for i := 0; i < 10; i++ {
		mu.Lock()
		v := count
		v++
		//time.Sleep(3 * time.Millisecond)
		runtime.Gosched()
		count = v
		fmt.Println(s, i, "count: ", count)
		mu.Unlock()
	}
	wg.Done()
}
