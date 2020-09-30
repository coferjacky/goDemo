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

	var x uint8 = 1<<1 | 1<<5
	fmt.Printf("%08b\n", x)

	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)

	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	//查看给定二进制数的1在哪些位置
	for i := uint(0); i < 8; i++ { //0转换为无符号整形
		if x&(1<<i) != 0 { //00100010 & 00000100
			fmt.Println(i) //在1和5的位置
		}
	}
	var m int8 = -100
	fmt.Printf("%064b\n", m)
	var z int8 = -100 >> 2
	fmt.Printf("%08b\n", z)
	fmt.Println(z)

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n")

	s := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x\n", s)
}
