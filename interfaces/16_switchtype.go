package main

import "fmt"

//判断一个接口内保存或实现的类型
/*
switch 接口变量.(type){
	case 类型1：
		变量是类型1时处理
    case 类型2：
		变量是类型2的时候处理
	default:
		变量不是类型1和2时处理
}

*/
func print_Type(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	}
}

func main() {
	print_Type(1024)
	print_Type("pig")
	print_Type("true")
}
