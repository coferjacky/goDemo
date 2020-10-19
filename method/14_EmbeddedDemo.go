package main

import (
	"fmt"
	"image/color"
)

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point2.X) // "1"
	cp.Y = 2
	fmt.Println(cp.Y) // "2"

	red := color.RGBA{0, 0, 255, 255}
	blue := color.RGBA{255, 0, 0, 255}
	var p = ColoredPoint{Point2{1, 2}, red, &Point3{1, 1}}
	var q = ColoredPoint{Point2{6, 6}, blue, &Point3{2, 2}}
	fmt.Println(p.Distance1(q.Point2))

	p.ScaleBy1(2)
	q.ScaleBy1(2)
	fmt.Println(p.Distance1(q.Point2)) // "10"

}
