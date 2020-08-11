package main

import (
	"fmt"
	"unsafe"
)

type Person2 struct {
	name string
	sex  byte
	age  int
}

func main() {
	//初始化方式1
	var p1 *Person2 = &Person2{
		name: "name1",
		sex:  'f',
		age:  19,
	}
	fmt.Println("p1", p1)

	//初始化方式2
	var tmp Person2 = Person2{"jack", 'm', 20}
	var p2 *Person2
	p2 = &tmp
	fmt.Println("p2", p2)

	//初始化方式3
	p3 := new(Person2)
	p3.name = "cofer"
	fmt.Println("p3", p3)

	test2(p3)               //传递结构体变量指针，推荐使用这种方式，不要直接传结构体变量，*****此种方式使用频率非常高
	p4 := unsafe.Sizeof(p3) // 值为8  *在64位系统中大小是一致的，都是8字节。
	fmt.Println("调用函数后p3", p4)

}

func test2(p *Person2) {
	p.name = "lufei"
	p.age = 229
	p.sex = 'm'
}
