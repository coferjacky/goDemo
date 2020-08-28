package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//第一步：建立SOCKET
	conn, err := net.Dial("udp", "192.168.0.80:8003") //指定服务器Ip+Port创建通信套接字
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	//第二步：主动写数据给服务器
	for {
		time.Sleep(time.Second * 3)
		conn.Write([]byte("are you ready?"))

		buf := make([]byte, 4096)
		//第三步：接收服务器回发的数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn err", err)
		}
		fmt.Println("服务器回发", string(buf[:n]))
	}
}
