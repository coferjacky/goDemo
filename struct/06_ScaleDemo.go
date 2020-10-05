package main

import "fmt"

/*
该案例主要是说明：结构体可以作为函数的的参数和返回值
*/

type Point struct{ X, Y int }

func main() {
	p := Point{10, 20}

	fmt.Println(Scale(p, 2))

}
func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}
