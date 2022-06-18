package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0
	const gr = 20

	wg.Add(gr)
	for i := 0; i < gr; i++ {
		mu.Lock()
		go func() {
			v := counter
			runtime.Gosched() //yield
			v++
			counter = v
			fmt.Println("Counter:", counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Total Counter: ", counter)
}
