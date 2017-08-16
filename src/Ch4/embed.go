package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w1, w2 Wheel
	
	w1 = Wheel{ Circle{Point{0,0}, 10}, 25 }
	w2 = Wheel {
		Circle: Circle{
			Point: Point{
				X: 1,
				Y: 1,
			},
			Radius: 12,
		}, 
		Spokes: 66,
	}

	fmt.Printf("Ordered: %#v\n", w1)
	fmt.Printf("Field-Value pairs: %#v\n", w2)
}