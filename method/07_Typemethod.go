package main

import "fmt"

//基本类型添加方法
type My_Int int

//myint添加Is_Zero()方法
func (m My_Int) Is_Zero() bool {
	return m == 0
}

//myInt 增加Add()方法
func (m My_Int) Add(other int) int {
	return other + int(m)
}

func main() {
	var b My_Int
	fmt.Println(b.Is_Zero())
	b = 1
	fmt.Println(b.Add(2))
}
