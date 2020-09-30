package main

import (
	"bytes"
	"fmt"
)

//int数组转string打印 并加方括号
func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v) //将v字符串输出到 buf指针指向的位置
	}
	buf.WriteByte(']')
	return buf.String()
}
