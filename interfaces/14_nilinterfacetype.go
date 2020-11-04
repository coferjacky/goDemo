package main

//空接口类型 能保存所有值
import "fmt"

func main() {

	var any interface{}
	any = 1
	fmt.Println(any)
	//赋值无论是什么，其类型一定是interface{}
	any = "hello"
	fmt.Println(any)

	any = false
	fmt.Println(any)
	//保存到空接口的值，如果直接取出指定类型的值时，会发生编译错误

	var ace int = 1
	var ice interface{} = ace
	//var b int=i  编译错误
	var bce int = ice.(int)
	fmt.Println(bce)

	//空接口再保存不同值后可以和其他变量一样使用==进行比较操作
	var a1 interface{} = 100
	var b1 interface{} = "hi1"
	fmt.Println(a1 == b1)

	//当接口总保存有动态类型值史，如果比较运算发生错误
	var c interface{} = []int{10}
	var d interface{} = []int{20}
	fmt.Println(c == d)
	fmt.Println(d)

}
