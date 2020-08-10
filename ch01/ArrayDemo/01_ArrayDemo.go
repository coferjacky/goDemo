package main

import "fmt"

//定义数组的

func main() {
	var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
	var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6

	e := a
	fmt.Println(e, b, c, d)

	var f = &b
	fmt.Println(b[0], b[1]) //数组变量访问
	fmt.Println(f[0], f[1]) //数组指针访问

	for i, v := range f { //通过数组指针迭代数组
		fmt.Println(i, v)
	}

}
