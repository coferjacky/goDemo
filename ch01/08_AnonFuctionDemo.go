package main

import (
	"fmt"
	"math"
)

func main() {
	//匿名函数使用方式一：定义时调用匿名函数
	func(data int) {
		fmt.Println("hello", data, data, data)
	}(100)
	//匿名函数使用方式二：匿名函数赋值给变量
	f := func(data string) {
		fmt.Println(data)
	}
	f("欢迎学习go语句")

	//匿名函数用作回调函数三：对每个元素进行求平方根操作
	arr := []float64{1, 9, 16, 25, 30}
	//数组元素开方
	visit(arr, func(v float64) { //第二个参数传入匿名函数，后面大括号实现了匿名的函数功能
		v = math.Sqrt(v)
		fmt.Printf("%.2f \n", v)
	})
	//数组元素平方
	visit(arr, func(v float64) {
		v = math.Pow(v, 2)
		fmt.Printf("%.0f \n", v)
	})

}

func visit(list []float64, f func(float64)) { //第二参数是函数变量
	for _, value := range list {
		f(value)
	}
}
