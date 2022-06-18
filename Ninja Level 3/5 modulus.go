package main

import "fmt"

const modulus int = 4

func main() {
	for num := 10; num <= 100; num++ {
		/*		if num%modulus == 0 {
					fmt.Println(num)
				}
		*/
		fmt.Printf("Num %v is divied by %v, and the remainder %v, and the modulus %v\n", num, modulus, num/modulus, num%modulus)
	}
}
