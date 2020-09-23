package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打印标准输入中多次出现的行，以重复次数开头
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
