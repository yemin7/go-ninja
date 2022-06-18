package main

import "fmt"

func main() {
	ninja := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	ninja = append(ninja, 52)
	fmt.Println(ninja)

	//ninja = append(ninja, []int{53, 54, 55}...)
	ninja = append(ninja, 53, 54, 55)
	fmt.Println(ninja)

	next_ninja := []int{56, 57, 58, 59, 60}
	ninja = append(ninja, next_ninja...)
	fmt.Println(ninja)

}
