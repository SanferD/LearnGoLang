package geometry

import "math"

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	var distance float64
	for i:=1; i!=len(path); i++ {
		distance += path[i-1].Distance(path[i])
	}
	return distance
}