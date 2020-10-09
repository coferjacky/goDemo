package main

import (
	"fmt"
	"os"
)

func main() {
	linenum, name := 12, "count"
	errorf(linenum, "undefined:%s", name)
}

//interface 表示函数后面最以后一个参数接受任意类型
func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
