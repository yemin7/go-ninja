package main

import "fmt"

func main() {
	ninja := make([]string, 10, 10)
	fmt.Println(ninja)
	fmt.Println(len(ninja))
	fmt.Println(cap(ninja))

	ninja = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii"}
	fmt.Println(ninja)
	fmt.Println(len(ninja))
	fmt.Println(cap(ninja))

	for v := 0; v < len(ninja); v++ {
		fmt.Print(ninja[v], "\t")
	}
}
