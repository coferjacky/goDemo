/*
将数组的单数和偶数分离出来并形成新数组,函数变量的示例
*/
package main

import (
	"fmt"
)

type processFunc func(int) bool //步骤1:定义了一个函数类型
func main() {
	slice := []int{4, 1, 2, 3, 4, 5, 6}
	fmt.Println("slice=", slice)
	odd := filter(slice, isOdd)
	fmt.Println("奇数元素", odd)
	even := filter(slice, isEven)
	fmt.Println("偶数元素", even)
}

//判断是否为偶数，是偶数则返回true
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

//判断是否为奇数，是偶数则返回false
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}
func filter(slice []int, f processFunc) []int {
	var result []int
	for _, value := range slice { //下划线“_”的意思就是忽略索引下标,返回value ,对于slice数组进行迭代
		fmt.Println(value)
		if f(value) {
			result = append(result, value) //将每次迭代的数据一个个加到数组里面
		}
	}
	return result
}
