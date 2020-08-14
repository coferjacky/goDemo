/*
字节读写文件 ***非常重要
案例：文件拷贝

*/
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//1打开读文件
	f_r, err := os.Open("c:/myjack.abc")
	if err != nil {
		fmt.Println("Open err", err)
		return
	}
	defer f_r.Close()
	//2 创建写文件
	f_w, err := os.Create("c:/myjack1.abc")
	if err != nil {
		fmt.Println("close err", err)
		return
	}
	defer f_w.Close()

	//3 从读文件中获取数据 放在缓存区
	//3.1 创建缓存区
	buf := make([]byte, 4096)
	//循环从读文件中获取数据，”原封不动“写到写文件中
	for {
		n, err := f_r.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Printf("读完毕 n= %d \n", n)
			return
		}
		f_w.Write(buf[:n]) //读多少写多少
	}

}
