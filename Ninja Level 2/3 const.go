package main

import "fmt"

const (
	a_untype = 25
	b_untype = "Onnut"
	c_untype = 25.52
)

const (
	a_type uint    = 7
	b_type string  = "UDelight"
	c_type float32 = 7.7
)

func main() {

	fmt.Printf("Untype values: %v\t%v\t%v\n", a_untype, b_untype, c_untype)
	fmt.Printf("Type values: %v\t%v\t%v", a_type, b_type, c_type)

}
