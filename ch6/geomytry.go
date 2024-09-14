package main

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot((p.X - q.X), (p.Y - q.Y))
}
func (p Point) Distance(q Point) float64 {
	return math.Hypot((p.X - q.X), (p.Y - q.Y))
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}


type Path []Point

func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}
