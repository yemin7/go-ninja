package main

import (
	"fmt"
	"sort"
)

func main() {
	slice_int := []int{2, 9, 10, 59, 93, 85, 60, 1, 15, 48}
	slice_string := []string{"Momo", "Dahyun", "Nancy", "Mina", "Sana"}

	sort.Ints(slice_int)
	sort.Strings(slice_string)

	fmt.Println(slice_int)
	fmt.Println(slice_string)
}
