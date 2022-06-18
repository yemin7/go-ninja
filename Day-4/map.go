package main

import "fmt"

func main() {
	t_map := map[string]int{
		"Mina": 26,
		"Tofu": 24,
	}

	fmt.Println(t_map)
	fmt.Println(t_map["Tofu"])
	v, ok := t_map["Mina"]
	fmt.Println(v, ok)

	if v, ok := t_map["Mina"]; ok {
		fmt.Println("Mina is", v)
	}
}
