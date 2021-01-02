package main

import "fmt"

func main() {
	s1 := "a/b$c{{}[]d"
	b1 := []byte(s1)
	fmt.Println(b1)      // [97 47 98 36 99 123 123 125 91 93 100]
	fmt.Println(len(b1)) //11
	fmt.Println(b1[:2])  //98
	s2 := "曹圆圆"
	b2 := []byte(s2)
	fmt.Println(b2) // [228 184 173 230 150 135], unicode，每个中文字符会由三个byte组成
	fmt.Println(len(b2))
	r1 := []rune(s1)
	fmt.Println(r1) // [97 98 99 100], 每个字一个数值
	r2 := []rune(s2)
	fmt.Println(r2) // [20013 25991], 每个字一个数值
}
