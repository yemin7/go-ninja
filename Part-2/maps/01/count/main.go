package main

import (
	"fmt"
	"log"

	"github.com/yemin7/go-ninja/Part-2/maps/01/datafile"
)

func main() {
	lines, err := datafile.GetStrings("strings.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
