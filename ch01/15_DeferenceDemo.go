/*
函数传int型参数
*/

package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("1.变量a的内存地址：%p,值为：%v \n \n", &a, a)
	fmt.Printf("========int型变量a的内存地址：%p \n \n", a)
	changeIntVal(a)
	fmt.Printf("2.变量a调用之后的内存地址：%p,值为：%v \n \n", &a, a) //调用函数之后a的地址
	changeIntPtr(&a)
	fmt.Printf("3.变量a调用之后的内存地址：%p,值为：%v \n \n", &a, a)

}
func changeIntVal(a int) {
	fmt.Printf("--------changeIntVal函数内：值参数a的内存地址：%p,值为：%v \n", &a, a)
	a = 90
}
func changeIntPtr(a *int) {
	fmt.Printf("--------changeIntPtr函数内：值参数a的内存地址：%p,值为：%v \n", &a, a)
	*a = 50
}
