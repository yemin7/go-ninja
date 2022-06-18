package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("custom error message %v", ce.info)
}

func main() {
	c1 := customErr{"custom info msg"}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
