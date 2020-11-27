package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果name没有指定的后缀名则上，否则就按照原来的名字来返回
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main() {
	//测试makesuffix的使用
	f := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f("winter"))
	fmt.Println("文件名处理后=", f("abc.jpg"))
	fmt.Println("文件名处理后=", f("abcd.abc"))
}
