package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("服务器开始监听了")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端来连接...")
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("accept() err=", err)

		} else {
			fmt.Printf("accept() succeed=%v\n,remoteaddress=%v\n", accept, accept.RemoteAddr())
		}
		go process(accept)

	}
	fmt.Printf("listen=%v", listen)
}

func process(conn net.Conn) {
	//这里我们循环接受客户端发送的数据
	defer conn.Close()
	for {
		//创建一个切片
		buf := make([]byte, 1024)
		fmt.Println("服务器等待客户端发送信息" + conn.RemoteAddr().String())
		//conn.Read(buf) 1、等待客户端通过conn发送信息 2、如果客户端没有write【发送】，那么协程阻塞这里
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("服务器端read err", err)
			return
		}
		//3显示客户端发送的内容到服务器
		fmt.Print(string(buf[:n]))
	}
}
