package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//获取服务器的应答包内容
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("httpget err", err)
	}
	defer resp.Body.Close() //resp.close 是返回一个是否关闭的状态

	//简单查看应答包
	fmt.Println("header:", resp.Header)
	fmt.Println("status:", resp.Status)

	buf := make([]byte, 4096)
	var result string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("-----Read finish")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("body read,err", err)
			return

		}
		result += string(buf[:n])

	}
	fmt.Printf("|%v|\n", result)

}
