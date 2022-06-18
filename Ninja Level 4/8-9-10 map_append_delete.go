package main

import "fmt"

func main() {
	ninja := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	/*	for k, v := range ninja {
			fmt.Printf("`%v`,", k)
			for _, v2 := range v {
				fmt.Printf("`%v`,", v2)
			}
			fmt.Println()
		}
	*/

	ninja["Dahyun"] = []string{"Tofu", "Twice"}
	for k, v := range ninja {
		fmt.Printf("`%v`,", k)
		for _, v2 := range v {
			fmt.Printf("`%v`,", v2)
		}
		fmt.Println()
	}

	fmt.Println()

	delete(ninja, "moneypenny_miss")
	for k, v := range ninja {
		fmt.Printf("`%v`,", k)
		for _, v2 := range v {
			fmt.Printf("`%v`,", v2)
		}
		fmt.Println()
	}
}
