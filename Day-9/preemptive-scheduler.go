package main

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		printHashes("A")
		wg.Done()
	}()

	go func() {
		printHashes("B")
		wg.Done()
	}()

	fmt.Println("Waiting to Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

func printHashes(s string) {
	for i := 0; i < 20; i++ {
		num := strconv.Itoa(i)
		sum := sha1.Sum([]byte(num))
		fmt.Printf("%s: %05d: %x\n", s, i, sum)
	}
	fmt.Println("Completed", s)
}
