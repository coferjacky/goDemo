package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("c:/myjack.abc", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("openFile err:", err)
		return
	}
	defer f.Close()
	fmt.Println("successful")
	//1 创建一个带缓冲区的reader（读写器）
	reader := bufio.NewReader(f)
	for {
		//2 从read缓冲区中，读取指定长度的数据。数据长度取决于   参数dlime,读到哪里结束，文件结束标记EOF要单独读一次获取到
		buf, err := reader.ReadBytes('\n') //读一行数  \n查看换行符
		if err != nil && err == io.EOF {
			fmt.Println("文件读取完毕")
			return
		} else if err != nil {
			fmt.Println("ReadBytes err:", err)
		}
		fmt.Print(string(buf))
	}
}
