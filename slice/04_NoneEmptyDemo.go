package main

import "fmt"

func main() {
	data := []string{"one", "", "three"} //初始化切片data
	data = nonempty(data)                //通常我们进行这个操作来解决下表的问题
	// fmt.Printf("%q\n",nonempty(data)) //共享了同一个数组，将该数组进行了覆盖，破坏了原来数组
	fmt.Printf("%q\n", data)
}
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
