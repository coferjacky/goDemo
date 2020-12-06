package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("c:/cofer.txt")
	if err != nil {
		fmt.Println("open file err", err)
	}
	//打印出的是个地址，所以file是指针
	fmt.Printf("file=%v", file)
	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file error", err)
	}
}
