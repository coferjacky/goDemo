package main

import "fmt"

/*
常用转义字符
*/
func main() {
	//  \t制表符
	fmt.Println("tom\tjack")
	//  \\转义为\
	fmt.Println("c:\\dfadsa")
	//  \r回车 从当前行的最前面进行输出，覆盖掉以前内容
	fmt.Println("你好啊曹菲\r他")
}
