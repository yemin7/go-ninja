package main

import (
	"fmt"
	"log"
)

func main() {

	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		ErrNorgateMath := fmt.Errorf("norgate math: square root of negative number")
		return 0, ErrNorgateMath
	}
	return 25, nil
}
