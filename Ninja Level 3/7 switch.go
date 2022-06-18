package main

import "fmt"

func main() {
	str1 := "Onnut"

	switch {
	case str1 == "UDelight":
		fmt.Println("It is UDelight.")

	case str1 == "OnNut":
		fmt.Println("It is OnNut.")

	case str1 == "Onnut":
		fmt.Println("It is Onnut.")
	}
}
