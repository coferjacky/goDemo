package main

import "fmt"

//一个被测试的函数
func addUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func getSub(n1 int, n2 int) int {
	return n1 - n2
}

func main() {
	ress := addUpper(10)
	fmt.Println(ress)
}
