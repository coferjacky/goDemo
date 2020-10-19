package main

import (
	"../packageDemo"
	"fmt"
)

//这里是方法值和方法表达式

func main() {
	p := packageDemo.Point{1, 2}
	q := packageDemo.Point{4, 6}
	//选择器返回了一个方法“值”：一个将方法(point.Distance)绑定到特定接受器变量的函数
	distanceFromP := p.Distance
	//可以不用指定接收器就被调用
	fmt.Println(distanceFromP(q))

	var origin packageDemo.Point
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy
	scaleP(2)  //{2,4}
	scaleP(4)  //}{8,16}
	scaleP(10) //{80,160}
	fmt.Println(p)
}
