package main

import "fmt"

//保持原来顺序的删除元素
func main() {
	s := []int{5, 6, 7, 8, 9, 10}
	fmt.Println(remove(s, 2)) //删除索引号为2的元素 7
}

func remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}
