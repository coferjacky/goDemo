package main

import (
	"../packageDemo"
	"fmt"
)

//这里是方法表达式

func main() {
	p := packageDemo.Point{1, 2}
	q := packageDemo.Point{4, 6}
	//方法表达式返回一个函数“值”,这里的dDistance 实际就是指定Point对象为接收器的一个方法
	distanceFromP := packageDemo.Point.Distance
	//可以不用指定接收器就被调用
	fmt.Println(distanceFromP(p, q))
	fmt.Printf("%T\n", distanceFromP)
	scale := (*packageDemo.Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)

}
