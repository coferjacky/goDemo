package main

import (
	"image/color"
	"math"
)

type Point2 struct{ X, Y float64 }
type Point3 struct{ X1, Y1 float64 }

type ColoredPoint struct {
	Point2
	Color color.RGBA
	*Point3
}

//接收器是Point类型,计算两点间矢量距离
func (p ColoredPoint) Distance1(q Point2) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//接收器是Point 指针类型
func (p *ColoredPoint) ScaleBy1(factor float64) {
	p.X *= factor
	p.Y *= factor
}
