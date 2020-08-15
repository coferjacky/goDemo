/*
指针的指针
*/

package main

import "fmt"

const COUNT int = 4

func main() {
	a := [COUNT]string{"abc", "ABC", "ddd", "eee"}
	i := 0
	//定义了指针数组
	var ptr [COUNT]*string
	fmt.Printf("%T,%v \n", ptr, ptr)
	for i = 0; i < COUNT; i++ {
		ptr[i] = &a[i]
	}
	fmt.Printf("%T,%v \n", ptr, ptr)

	//获取指针数组中第一个值，其实就是一个地址
	fmt.Println(ptr[1])
	for i = 0; i < COUNT; i++ {
		fmt.Printf("a[%d]=%s \n", i, *ptr[i])
	}
}
