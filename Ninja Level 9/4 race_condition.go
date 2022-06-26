package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	//var counter int
	counter := 0
	const gr = 20

	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			mu.Lock()
			v := counter
			//runtime.Gosched() //yield
			v++
			//log.Println("logging")
			counter = v
			fmt.Println("Counter:", counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Total Counter: ", counter)
}
