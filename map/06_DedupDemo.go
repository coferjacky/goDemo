package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // 遇到 \n 或者\r\n循环一次，每输入一行内容必须回车，退出循环按组合键 Ctrl +D
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup:%v\n", err)
		os.Exit(1)
	}
}
