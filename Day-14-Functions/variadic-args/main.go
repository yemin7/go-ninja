package main

import "fmt"

func main() {
	data := []float64{34, 52, 63, 91, 18}
	fmt.Println(average(data...))
}

func average(arg ...float64) float64 {
	var total float64
	for _, v := range arg {
		total += v
	}
	return (total / float64(len(arg)))
}
