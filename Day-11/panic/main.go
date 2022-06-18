package main

import "os"

func main() {

	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
	//Panic function call panic after writing the log message.
}
