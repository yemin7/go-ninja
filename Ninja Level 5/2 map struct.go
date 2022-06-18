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

	ninja := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range ninja {
		fmt.Println("First Name:", v.first)
		fmt.Println("Last Name:", v.last)
		fmt.Print("Ice cream flavors: ")
		for _, ic := range v.flavor {
			fmt.Print(ic, "\t")
		}
		fmt.Println("\n---------")
	}
}
