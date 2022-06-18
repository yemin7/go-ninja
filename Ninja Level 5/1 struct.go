package main

import "fmt"

type person struct {
	first, last string
	flavor      []string
}

func main() {
	p1 := person{
		first: "みな",
		last:  "名井",
		flavor: []string{
			"chocolate",
			"strawberry",
			"milk",
		},
	}
	p2 := person{
		first: "さな",
		last:  "湊崎",
		flavor: []string{
			"chocolate",
			"strawberry",
			"milk",
		},
	}
	fmt.Println(p1.first, p1.last)
	for _, v := range p1.flavor {
		fmt.Print(v, "\t")
	}
	fmt.Println()
	fmt.Println(p2.first, p2.last)
	for _, v := range p2.flavor {
		fmt.Print(v, "\t")
	}
}
