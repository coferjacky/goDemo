package main

import (
	"fmt"
)

func main() {
	//位运算左移n 相当于数字*2的n次方;位右移n 相当于数字/2的n次方
	a := 3
	fmt.Println(a << 4)
	b := 32
	fmt.Println(b >> 4)
	/*
		异或   每一位上只要两个输入数的同一位置不同则为1
	*/
	d := 20
	e := 5
	fmt.Println(d ^ e)

}
