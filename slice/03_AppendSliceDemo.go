package main

import "fmt"

func main() {
	var runes []rune //rune是int32的别名，用来处理utf-9字符的
	for _, r := range "Hello，世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	arrs := [5]int{1, 2, 3, 4, 5}
	s := arrs[0:5]
	fmt.Println("原数组为：", arrs)
	fmt.Println("输入切片：", s)
	ints := appendInt(s, 33)
	fmt.Println("运算完毕后得数组为：", arrs)
	fmt.Println("返回切片：", ints)
	ints1 := ints[:]
	fmt.Println("返回切片：", ints1)

}

/*
每次调用appendInt函数，必须先检测：
slice底层是否有足够的容量保存新添元素，如果有足够的容量，
直接扩展slice，返回的slice与输入的slice共享底层数组

如果没有足够的容量，函数会先分配一个足够的大的slice用于保存新的结果，先将输入的切片x复制到新空间，
然后添加y元素。结果z和输入x引用了不同的底层数组




*/
func appendInt(x []int, y int) []int { //只有接受切片
	var z []int //声明一个切片z

	zlen := len(x) + 1  //切片z的长度是参数长度+1，比如传入s 则此处为4
	if zlen <= cap(x) { //此时的长度小于容量
		z = x[:zlen] //此时的z切片就是x[0:4]  匿名数组为 0 0 1
	} else {
		zcap := zlen         //5+1=6 则将6赋给zcap
		if zcap < 2*len(x) { //如果 6<2*6 则直接将其容量扩展1倍
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap) //新建一个匿名数组，容量为12
		copy(z, x)                  //将x复制给z
	}
	z[len(x)] = y
	return z
}
