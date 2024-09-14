package main

import (
	"fmt"
	"image/color"
)

type ColoredPoint struct  {
	*Point 
	Color color.RGBA
}
func main() {
	red := color.RGBA{255, 0, 0, 255}
	cp := ColoredPoint{&Point{1, 1}, red}
	fmt.Println(cp.Distance(Point{2, 2})) // "1.4142135623730951")
	cp.ScaleBy(2)
	fmt.Println(cp.Point) // "{2 2}"
}