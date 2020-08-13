package main

import (
	"fmt"

	"os"
)

func main() {
	// 创建文件
	f, err := os.Create("c:/myjack.abc")
	if err != nil {
		fmt.Println("create err:", err)
		return
	}
	defer f.Close()
	fmt.Println("success")

	//只读方式进行打开文件（只读方式打开的文件） 打开有三个参数 文件路径含文件名、 文件打开权限、 一般是数字6  三个参数
	f1, err1 := os.Open("c:/myjack.abc")
	if err1 != nil {
		fmt.Println("open err:", err1)
		return
	}
	defer f1.Close()
	fmt.Println("successopen")

	//往文件写入数据 (只读报错)
	_, err3 := f1.WriteString("********88888")
	if err3 != nil {
		fmt.Println("write err", err3)
		return
	}
	fmt.Println("successful")

}
