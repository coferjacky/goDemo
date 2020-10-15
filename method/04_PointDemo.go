package main

import "fmt"

//非指针类型的接收器
type Point struct {
	X int
	Y int
}

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
