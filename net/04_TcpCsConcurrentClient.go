package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.0.101:8001")
	if err != nil {
		fmt.Println("conn err", err)
		return
	}
	defer conn.Close()
	go func() {
		str := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.stdinlreadd err", err)
				continue
			}
			conn.Write(str[:n])
		}

	}()
	//回显服务器回发的大写数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("检查服务器关闭，则客户端也关闭")
			return
		}
		if err != nil {
			fmt.Println("conn.read err", err)
			return
		}
		fmt.Println("客户都读到服务器的回发", string(buf[:n]))

	}
}
