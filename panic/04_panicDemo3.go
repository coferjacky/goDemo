package main

import "fmt"

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("发生故障", err)
			fmt.Println("发送邮件给管理员")
		}
	}()
	num1 := 10
	num2 := 0
	num3 := num1 / num2
	fmt.Println(num3)
}
func main() {
	test()
	fmt.Println("正常执行")
}
