package main

import (
	"fmt"
	"net"
	"os"
)

func errFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		os.Exit(1) //将当前进程结束
	}

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	errFunc(err, "net.Listener err:")
	defer listener.Close()

	conn, err := listener.Accept()
	errFunc(err, "conn err:")
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	errFunc(err, "conn.read")
	fmt.Println(string(buf[:n]))
}
