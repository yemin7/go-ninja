package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	//counter := 0
	var counter int64
	const gr = 20

	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			//	v := counter
			atomic.AddInt64(&counter, 1)
			//runtime.Gosched() //yield
			//v++
			//counter = v
			//fmt.Println("Counter:", counter)
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Total Counter: ", counter)
}
