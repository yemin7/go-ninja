package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter int
	const gr = 20

	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			v := counter
			//runtime.Gosched() //yield
			v++
			counter = v
			fmt.Println("Counter:", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Total Counter: ", counter)
}
