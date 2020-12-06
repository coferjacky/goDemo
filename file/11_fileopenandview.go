package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("c:/myjack.abc")
	if err != nil {
		fmt.Println("open file err")
		return
	}
	defer file.Close()
	//创建一个*reader 是带缓冲的 默认4096个字节
	reader := bufio.NewReader(file)
	//循环读取文件内容
	for {
		//读到换行符号就结束
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("读取完毕了")
}
