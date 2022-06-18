package main

import "fmt"

const pi float32 = 3.142

type circle struct {
	radius int
}

type square struct {
	side int
}

//method for circle
func (c circle) area() float32 {
	//	area := pi * float32(c.radius*c.radius)
	return pi * float32(c.radius*c.radius)
	//	fmt.Println("Circle area: ", area)
}

//method for square
func (s square) area() float32 {
	return float32(s.side * s.side)
	//	square := float32(s.side * s.side)
	//	fmt.Println("Square area: ", square)
}

type c_shape interface {
	area() float32
}

func info(ch c_shape) {
	c := ch.area()
	fmt.Println(c)
}

func main() {
	c1 := circle{
		radius: 3,
	}
	s1 := square{
		side: 3,
	}
	c1.area()
	s1.area()

	info(c1)
	info(s1)
}
