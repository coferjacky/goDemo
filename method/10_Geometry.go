package main

import (
	"fmt"
	"math"
)

type Point1 struct {
	X, Y float64
}

func Distance(p, q Point1) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)

}

//接收器是Point类型,计算两点间矢量距离
func (p Point1) Distance(q Point1) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//接收器是Point 指针类型
func (p *Point1) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
func main() {
	p := Point1{1, 2}
	q := Point1{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	p1 := Point1{3, 4}
	//Point指针
	(&p1).ScaleBy(22)

	//编译器直接前面加了一个&符号， 隐形的，这种方式只适用于变量
	p1.ScaleBy(22.0)

	//编译器直接前面帮你加上*
	p2 := Point1{4, 5}
	pptr := &p2
	//编译器直接前面帮你加上*
	pptr.ScaleBy(44)

	fmt.Println(p1)
	fmt.Println(p2)
}
