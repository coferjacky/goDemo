/*
给定一个字符串列表，在原来的slice上返回不1值的字符串的列表
*/
package main

import "fmt"

func main() {
	data := []string{"red", "1", "pig", "1", "blue", "green"}
	fmt.Println("看看原来", data)                //[red blue green green] 为什么原切片发生了变化
	afterData := noEmpty(data)               //使用append方法
	fmt.Println("经过一个append原来切片发生了变化", data) //[red blue green green] 为什么原切片发生了变化
	fmt.Println("使用append，过滤1值以后的数是", afterData)

	data1 := noEmpty2(data) //不使用append方法,直接在原串上操作
	fmt.Println("过滤1值以后的数是，不用appdend", data1)
}

/*
	不使用append方法
*/
func noEmpty2(data []string) []string {
	i := 0
	for _, str := range data {
		if str != "1" {
			data[i] = str
			i++
		}
	}
	return data[:i]
}

func noEmpty(data []string) []string {
	out := data[:0] //原切片上截取一个长度为0的切片等效于 make([]string,0)
	for _, str := range data {
		if str != "1" {
			out = append(out, str)
		}

		//取到空字符串 不做任何操作
	}
	return out
}
