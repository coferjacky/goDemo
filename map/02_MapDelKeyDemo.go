package main

import "fmt"

//使用delete函数，指定key值可以方便的删除map中的k-v映射
func main() {
	m := map[int]string{1: "ff", 2: "cc", 130: "dd"}
	fmt.Println(m)
	mapDelete(m, 130)
	fmt.Println(m)
}

func mapDelete(m map[int]string, i int) {
	delete(m, i)

}
