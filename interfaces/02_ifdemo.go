package main

import "fmt"

//定义一个接口，两个方法
type Testinterface interface {
	TestString() string
	TestInt() int
}

//定义一个结构体
type TestMethod struct {
	name string
	age  int
}

//结构体两个方法隐式实现
func (t *TestMethod) TestString() string {
	return t.name
}

func (t *TestMethod) TestInt() int {
	return t.age
}

func main() {
	T1 := &TestMethod{"jack", 12} //TestMethod类的指针类型实例变量
	T2 := TestMethod{"gos", 13}   //TestMethod类的值类型实例变量
	fmt.Println("T1:", T1)
	fmt.Println("T2:", T2)

	var Test1 Testinterface //声明一个接口类型变量，接口只能是
	Test1 = T1
	fmt.Println(Test1.TestString())
	fmt.Println(Test1.TestInt())

	Test1 = T2
	fmt.Println(T2.TestInt())
	fmt.Println(T2.TestString())

}
