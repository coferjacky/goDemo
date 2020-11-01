package main

import (
	"fmt"
	"sort"
)

//利用包总的String_Slice快速完成排序,不用像06这么麻烦了要自己写实现-字符串切片的便捷排序
func main() {
	namess := sort.StringSlice{
		"3.fdadsaf",
		"5.fdafdsafsa",
		"zfdafasdf",
		"2fdafdasfa",
		"1fdasfdsaf",
	}
	sort.Sort(namess)

	sort.Sort(namess)
	for _, v := range namess {
		fmt.Printf("%s\n", v)
	}
}
