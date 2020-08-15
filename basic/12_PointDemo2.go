/*
基本数据类型指针作为函数的参数
*/
package main

import (
	"fmt"
)

func main() {
	a := 58
	fmt.Println("函数调用之前a的值：", a)
	fmt.Printf("%T \n", a)
	fmt.Printf("%x \n", &a)
	var b *int = &a
	change(b)
	fmt.Println("函数调用之后的a的值", a)
}
func change(vul *int) {
	*vul = 15
}
