package main

/*
封装wcFunc()函数，接收一段英文字符串str，返回一个map，记录str中每个“词”出现次数
*/
import (
	"fmt"
	"strings"
)

func main() {
	str := "i love my work and i love my family too"
	fmt.Println(str)
	mRet := wordCountFun(str)

	for k, v := range mRet {
		fmt.Println("key=", k, " value=", v)
	}
}

func wordCountFun(str string) map[string]int {
	s := strings.Fields(str)  //将字符串拆分为字符串切片
	m := make(map[string]int) //创建一个用于存储word出现次数的map key是string value是 次数

	for i := 0; i < len(s); i++ { //遍历拆分后的字符串切片,并逐一将字符串赋放入map的key中
		if _, has := m[s[i]]; has {
			m[s[i]] = m[s[i]] + 1
		} else {
			m[s[i]] = 1
		}
	}
	return m
}
