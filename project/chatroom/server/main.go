package main

import (
	"fmt"
	"net"
)

//处理和客户端的通讯
//获取套接字
func process(conn net.Conn) {

}

func main() {
	fmt.Println("服务器在8889端口监听")
	listen := net.Listener("tcp", "0.0.0.0:8889")

	//监听成功就等待客户端来连接服务器
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
