package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//1、首先将原文件内容读取内存
	//2、将读取到内容写到新文件里面
	filePathSource := "c:/myjack.abc"
	filePathTarget := "c:/myjackba.abc"
	content, err := ioutil.ReadFile(filePathSource)
	if err != nil {
		//说明读文件错误
		fmt.Printf("read file err=%v", err)
		return
	}
	err = ioutil.WriteFile(filePathTarget, content, 0666)
	if err != nil {
		fmt.Printf("write err=%v", err)
	}
}
