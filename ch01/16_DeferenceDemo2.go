/*
函数传slice型参数
*/
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Printf("1.变量a的内存地址是：%p,值为%v \n \n", &a, a)
	fmt.Printf("切片型变量a的值是地址，该地址是：%p \n \n", a)
	changeSliceVal(a) //传值
	fmt.Printf("2.变量a的内存地址是：%p,值为%v \n \n", &a, a)
	changeSlicePtr(&a) //传引用
	fmt.Printf("3.changeSlicePtr:函数调用后，变量a的内存地址是:%p,值是%v", &a, a)

}

//将实参中值（地址）复制给形参，复制完毕，此时就有两个指针指向了同一个数组区域
func changeSliceVal(a []int) {
	fmt.Printf("-----changeSliceVal函数内:值参数a的内存地址是:%p,值为%v \n", &a, a)
	fmt.Printf("-----changeSlicePtr函数内：值参数a的内存地址是：%p \n", a)
	a[0] = 99
}
func changeSlicePtr(a *[]int) {
	fmt.Printf("-------changeSlidePtr函数内：指针参数a的内存地址是%p,值为%v \n", &a, a)
	(*a)[1] = 250
}
