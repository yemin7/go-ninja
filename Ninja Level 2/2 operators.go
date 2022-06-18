package main

import "fmt"

const (
	num1 = 7
	num2 = 25
)

func main() {
	a := (num1 == num2)
	b := (num1 <= num2)
	c := (num1 >= num2)
	d := (num1 != num2)
	e := (num1 < num2)
	f := (num1 > num2)

	fmt.Println(a, b, c, d, e, f)
}
