package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	name string
	sex  byte
	age  int
}

func test(man Person) {
	fmt.Println("test temp size:", unsafe.Sizeof(man))
	man.name = "name1"
	man.age = 33
	fmt.Println("调用的man:", man)
}
func main() {
	//1 顺序初始化 必须全部初始化完整
	var man Person = Person{
		name: "caofie",
		sex:  'm',
		age:  21,
	}
	fmt.Println("man:", man)

	//2 部分初始化
	man2 := Person{
		name: "caoyy",
		age:  13,
	}
	fmt.Println("man2:", man2)
	fmt.Println("man2.name:", man2.name)

	//3 索引成员初始化
	var man3 Person
	man3.name = "yuanj"
	man3.sex = 'o'
	man3.age = 16
	fmt.Println("man3:", man3)
	// 如果要修改成员值
	man3.age = 19
	fmt.Println("man3:", man3)

	//4 相同结构体赋值
	var temp Person
	temp = man3
	temp.age = 200
	fmt.Println("temp:", temp)
	fmt.Println("man3:", man3)

	//5 函数内部使用结构体参数传参
	var tmp Person
	fmt.Println("main temp size", unsafe.Sizeof(tmp))
	test(tmp)
	fmt.Println("tmp:", tmp)

	fmt.Printf("&tmp=%p \n", &tmp)
	fmt.Printf("&tmp.name=%p \n", &temp.name)
	fmt.Println("main temp size:", unsafe.Sizeof(true))

}
