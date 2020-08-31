package main

import "fmt"

//1声明
var a int
var b string
var c []float32
var d func() bool //声明一个返回值为波尔类型的函数变量，这种形式为回调函数，函数以变量的形式保存下来，需要时重新调用这个函数
var e struct {
	x int
}

//1.1声明批量格式
var (
	f int
	g string
	h []float32
	i func() bool
	j struct {
		x int
	}
)

//2声明同时初始化
//2.1一般声明并初始化
var hp int = 100

//2.2编译器来推导类型格式
var hp1 = 100

//hpp:=2222 此处编译错误，因为全局变量声明必须以 var 关键字开头，如果想要在外部包中使用全局变量的首字母必须大写

func main() {
	//2.3 短变量声明（局部变量可以采用此模式进行）
	hpp := 222 //之前如果声明过则会编译会发生错误
	fmt.Println("HP3:=", hpp)
}

/*
声明的格式
1、标准格式 : var  变量名  变量类型
2、批量格式：
var (
	a int
    b string
    c []f

)
3、整数和浮点类型变量默认值 0
   字符串变量默认值为空字符串   null用于指示指针不引用有效对象针对指针不占用空间
   布尔变量默认值为bool
	切片 函数 指针变量的默认为nil  nil表示无值，针对对象占用空间

声明并完成初始化
1、标准格式 : var  变量名  变量类型=表达式

*/
