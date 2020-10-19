package main

import "fmt"

//非指针类型的接收器
type Point struct {
	X int
	Y int
}

//定义一个切片
type Path []Point

func main() {
	p1 := Point{1, 1}
	p2 := Point{2, 2}
	result := p1.Add(p2)
	fmt.Println(result)
	fmt.Println(p1) //p1并没有发生变化
}
func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

func (p Point) AddPoint(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}
func (p Point) SubPoint(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.AddPoint
	} else {
		op = Point.SubPoint
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}
