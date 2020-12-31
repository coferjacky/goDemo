package main

import (
	"fmt"
	"net"
)

//处理和客户端的通讯
//获取套接字
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()
	buf := make([]byte, 8096)
	//循环的客户端发送的信息

	for {
		fmt.Println("读取到了客户端发送的数据 ...")
		n, err := conn.Read(buf[:4])
		if err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		fmt.Println("读到的buf=", buf[0:4])
	}

}

func main() {
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	//监听成功就等待客户端来连接服务器
	if err != nil {
		fmt.Println("net.listen err=", err)
	}
	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listener.accept err", err)

		}
		//连接成功，则启动一个协程和客户端保持交互和通讯
		go process(conn)
	}
}
