/*
函数传结构体
*/
package main

import "fmt"

type Teacher struct {
	name    string
	age     int
	married bool
	sex     int8
}

func main() {
	a := Teacher{"Steven", 35, true, 1}
	fmt.Printf("1.变量a的内存地址是：%p,值为：%v \n \n", &a, a)
	fmt.Printf("struct型变量a内存地址是：%p \n \n", a)
	//传值
	changeStructVal(a)
	fmt.Printf("函数调用后，变量a的内存地址是：%p,值为：%v \n \n", &a, a)
	//传引用
	changeStructPtr(&a)
	fmt.Printf("函数调用后，变量a的内存地址是：%p,值为：%v \n \n", &a, a)

}

func changeStructPtr(a *Teacher) {
	fmt.Printf("------changeStructPtr函数内：指针参数a的内存地址是：%p,值为%v \n", &a, a)
	(*a).married = false
	(*a).age = 28
}

func changeStructVal(a Teacher) { //这里的改动不会影响原来的结构体
	fmt.Printf("-----changeStructVal函数内：值参数a的内存地址是：%p,值为%v \n", &a, a)
	fmt.Printf("-----changeStructVal函数内：值参数a的内存地址是：%p \n", a)
	a.name = "jack"
	a.age = 29
	a.married = false

}
