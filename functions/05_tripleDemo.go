package main

import "fmt"

//被延迟执行的匿名函数甚至可以修改函数返回给调用者返回值
func main() {
	fmt.Println(triple(4))
}
func triple(x int) (result int) {
	defer func() { result += x }() //将结果最后修改后返回
	return double1(x)
}
func double1(x int) int {
	return x + x
}
