/*
函数变量的例子
*/
package main

import (
	"fmt"
	"strings"
)

func main() {

	//处理字符串，索引为偶数的为大写，奇数为小写
	result := StringToLower("abCFFGDFgF5GDDDddduU", processCase)
	fmt.Println(result)
	result = StringToLower2("ABCCCdcdfdsEdd", processCase)
	fmt.Println(result)

	//将数组中的

}
func processCase(str string) string {
	result := ""
	//处理字符串，索引为偶数的为大写，奇数为小写
	for i, value := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}
func StringToLower(str string, f func(string) string) string { //注意：这里的返回值没有声明返回值变量，可以省略包含返回值的括号
	fmt.Printf("%T \n", f)
	return f(str)
}

type caseFunc func(string) string //声明一个函数类型，通过type关键字，caseFunc会形成一个新的类型
func StringToLower2(str string, f caseFunc) string {
	fmt.Printf("%T \n", f)
	return f(str)
}
