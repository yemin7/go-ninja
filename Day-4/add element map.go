package main

import "fmt"

func main() {
	t_map := map[string]int{
		"Mina": 26,
		"Tofu": 24,
	}
	fmt.Println(t_map)

	t_map["Sana"] = 26

	fmt.Println(t_map)

	for k, v := range t_map {
		fmt.Println(k, v)
	}
}
