package main

import (
	"fmt"
	"sort"
)

//接口实现不受限于结构体，任何类型都可以实现接口
//这里[]string是系统定制好的类型，无法让这个类型去实现sort.Interface排序接口。因此需要将[]string定义为自定义类型
type My_String_List []string

//实现获取元素数量的len()方法，返回字符串切片元素数量
func (m My_String_List) Len() int {
	return len(m)
}

//实现sort.Interface接口的比较元素方法
func (m My_String_List) Less(i, j int) bool {
	return m[i] < m[j]
}

//实现sort.Interface接口的交换元素方法
func (m My_String_List) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	names := My_String_List{
		"3.fdadsaf",
		"5.fdafdsafsa",
		"zfdafasdf",
		"2fdafdasfa",
		"1fdasfdsaf",
	}
	fmt.Printf("%T \n", names)
	sort.Sort(names)
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}
