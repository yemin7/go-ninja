package main

import "fmt"

func main() {
	t_map := map[string]int{
		"Mina": 26,
		"Tofu": 24,
	}

	t_map["Sana"] = 26
	fmt.Println(t_map)

	delete(t_map, "Sana")
	fmt.Println(t_map)

}
