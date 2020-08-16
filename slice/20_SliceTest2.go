/*
	字符去重，通过切片去重就是在原来的数组上将
*/
package main

import "fmt"

func main() {
	data := []string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	fmt.Println("原切片", data)
	afterData := noSame(data)
	fmt.Println("经过nosame函数去重以后的原切片", data)
	fmt.Println("经过nosame函数去重以后的新切片afterData", afterData)
}

func noSame(data []string) []string {
	out := data[:1] //这里我相当于初始化一个切片，然后切片里面有data中的第一个数 即red,其实还是指向原来的切片元素的

	for _, word := range data {
		//比较取出的word是否在
		i := 0
		for ; i < len(out); i++ {
			if out[i] == word {
				break
			}
		}
		if i == len(out) {
			out = append(out, word)
		}
	}
	return out
}
