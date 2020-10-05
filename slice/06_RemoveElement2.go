package main

import "fmt"

//不保持原来顺序的删除元素
//实际上就是在操作数组，将数组最后一个数字覆盖到索引i的元素，然后取切片返回，数组的长度和元素数并未发生变化
func main() {
	s := [...]int{5, 6, 7, 8, 9}
	sl := s[:]
	fmt.Println(remove2(sl, 2))
	fmt.Println(s)
}
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
