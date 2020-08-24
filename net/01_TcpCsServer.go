package main

//创建用于监听的socket
import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen er:", err)
		return
	}
	defer listen.Close()
	fmt.Println("服务器等待客户端建立连接.....：")
	//阻塞监听客户端连接请求，成功建立连接，返回用于通信的socket
	conn, err := listen.Accept() //阻塞监听客户端连接请求
	if err != nil {
		fmt.Println("listen.Accept() err")
		return
	}
	defer conn.Close()
	fmt.Println("服务器与客户端连接建立")
	//读取客户端发送的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn read err:", err)
		return
	}
	conn.Write(buf[:n]) //读多少写多少，原封不动
	//处理数据
	fmt.Println("服务器读到数据,", string(buf[:n]))

}
