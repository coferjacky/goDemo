package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//组织一个udp的地址结构,这是服务器的Ip+端口
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("resolveUpdAddr err:", err)
		return
	}
	//创建用户通信的socket
	udpConn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Println("listenudp err:", err)
		return
	}
	defer udpConn.Close()

	//读取客户端发送的数据
	buf := make([]byte, 4096)
	for {
		//ReadFromUDP返回拷贝字节数和数据包来源地址。
		n, cltAddr, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("reaudperr", err)
			return
		}
		//模拟数据的处理
		fmt.Printf("服务器读到%v数据:%s\n：", cltAddr, string(buf[:n]))

		//回写数据给客户端
		daytime := time.Now().String()
		_, err = udpConn.WriteToUDP([]byte(daytime), cltAddr)
		if err != nil {
			fmt.Println("write error", err)
			return
		}
	}

}
