package main

import "fmt"

func main() {
	str1 := "Swimming"

	switch str1 {
	case "FootBall":
		fmt.Println("Fav sport is football.")
	case "Tennis":
		fmt.Println("Fav sport is Tennis.")
	case "Swimming":
		fmt.Println("Fav sport is ", str1)
	}
}
