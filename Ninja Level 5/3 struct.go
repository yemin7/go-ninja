package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	truck1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: false,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "yellow",
		},
		luxury: true,
	}

	fmt.Printf("Doors: %v, Color: %v, fourWheel: %v\n", truck1.doors, truck1.color, truck1.fourWheel)
	fmt.Printf("Doors: %v, Color: %v, Luxury: %v\n", sedan1.doors, sedan1.color, sedan1.luxury)
}
