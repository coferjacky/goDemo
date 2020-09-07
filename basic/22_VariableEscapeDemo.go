package main

import "fmt"

func main() {
	//声明a变量并打印
	var a int
	//调用void函数
	void()
	fmt.Println(a, dummy(0))
}
func void() {

}
func dummy(b int) int {
	var c int
	c = b
	return c
}
