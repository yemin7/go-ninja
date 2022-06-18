package main

import "fmt"

func main() {
	str1 := []string{"Jame", "Bond", "Shaken not stiired"}
	str2 := []string{"Miss", "Moneypenny", "Hello, James"}

	ninja := [][]string{str1, str2}

	for _, v := range ninja {
		fmt.Println(v)
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
}
