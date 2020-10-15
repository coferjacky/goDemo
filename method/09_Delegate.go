package main

import "fmt"

//声明一个结构体
type class struct {
}

//结构体添加do方法
func (c *class) Do(v int) {
	fmt.Println("call method do:", v)
}

//普通函数的do方法
func func_Do(v int) {
	fmt.Println("call function do:", v)
}

//无论是普通函数还是结构体方法，只要他们签名一样（所谓签名一样就是说结构体方法的参数和普通函数的参数一致），与他们签名一致的函数变量就可以保存普通函数或是结构体方法
func main() {
	//声明一个函数回调
	var delegate func(int) //这里的参数必须和普通函数的参数和结构体方法的参数一致
	c := new(class)
	//回调设置为c的do方法
	delegate = c.Do
	//调用
	delegate(100)
	//将回调设置为普通函数
	delegate = func_Do
	//调用
	delegate(100)
}
