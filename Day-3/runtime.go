package main

import (
	"fmt"
	//	"runtime"
)

func main() {
	//	fmt.Println(runtime.GOOS)
	//	fmt.Println(runtime.GOARCH)

	string1 := "Hello"
	fmt.Printf("%v\t%T\n", string1, string1)

	byte_string := []byte(string1) //uint8
	//	byte_string := []rune(string1) //int32
	fmt.Println(byte_string)
	fmt.Printf("%T", byte_string)
}
