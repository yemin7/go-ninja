package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

func main() {

	truck := struct {
		vehicle
		fourWheel bool
	}{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: false,
	}

	fmt.Println(truck)

}
