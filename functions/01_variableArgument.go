package main

import "fmt"

/*
sum函数返回任意个int形参数的和
调用者隐式创建一个数组，并将原始参数复制到数组中，再把数组的一个切片作为参数传给被调函数
*/
func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	//如果原始参数已经是一个切片,则变量后面加省略符即可
	values := []int{1, 3, 5, 7}
	fmt.Println(sum(values...))
}

//调用者隐式创建一个数组，并将原始参数复制到数组中，再把数组的一个切片作为参数传给被调函数
func sum(vals ...int) int {

	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
