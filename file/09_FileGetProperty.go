package main

import (
	"fmt"
	"os"
)

/*
命令行参数 go run xxx.go arg1 arg2 arg3 arg4
xxx.go 第0个参数
arg1  第1个参数
.....

*/
func main() {
	list := os.Args //获取命令行参数

	fmt.Print(list)
	fmt.Println("index is 3 arg:", list[1])
	//go run 09_FileGetProperty.go aa bb cc 则取到cc参数了
	if len(list) != 2 {
		fmt.Println("格式为：go run xxx.go 文件名")
		return
	}
	//提取文件名
	path := list[1]
	//获取文件属性,注意 path为完整路径
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("path error", err)
		return
	}
	fmt.Println("文件名：", fileInfo.Name())
	fmt.Println("文件大小", fileInfo.Size())
}
