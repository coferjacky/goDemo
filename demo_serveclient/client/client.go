package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	dial, err := net.Dial("tcp", "192.168.100.169:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	//功能1：客户端可以发送单行数据然后退出
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readstring err=", err)
	}
	write, err := dial.Write([]byte(readString))
	if err != nil {
		fmt.Println("conn.write err=", err)
	}
	fmt.Printf("客户端发送了%d字节的数据,并退出", write)
}
