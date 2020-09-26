package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打印标准输入中多次出现的行，以重复次数开头
	counts := make(map[string]int)

	//定义变量bufio.Scanner类型变量 input
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { //每次调用input.Scan()，即读入下一行，并移除行末的换行符
		counts[input.Text()]++
		if input.Text() == "bye" {
			break
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
