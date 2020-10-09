package main

import "fmt"

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

//函数返回一个匿名函数值，这种方式定义就能访问完整的词法环境
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
