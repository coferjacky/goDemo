package main

import (
	"fmt"
	"strings"
)

//去掉 路径和后缀的信息
func main() {
	fmt.Println(basename("c:/121/222/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))

	fmt.Println(basename2("c:/121/222/c.go"))
	fmt.Println(basename2("c.d.go"))
	fmt.Println(basename2("abc"))
}

//手工硬编码模式
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

//stings.LastIndex函数实现功能
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") //从字符串最后字段来截取
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s

}
