package main

import "fmt"

//数字每隔三位取逗号
func main() {
	fmt.Println(comma("1001289212"))
}
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

//结果 1,001,289,212
