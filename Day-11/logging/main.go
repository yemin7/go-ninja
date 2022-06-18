package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
		//log.Fatalln(err)

		/* the Fatal function call
		os.Exit(1) after writing the log message*/

		//Fatalln is equivalent to Println() followed by a call to os.Exit(1).
	}
}
