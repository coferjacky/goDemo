package main

import (
	"fmt"
	"net"
	"os"
)

/*
模拟一下浏览器的行为，看看httpResponseDemo发来的包
*/

func errFunc2(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		os.Exit(1) //将当前进程结束
	}

}
func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	errFunc2(err, "dial err")
	defer conn.Close()

	httpRequest := "GET /itcast HTTP/1.1\r\nHost:www.baidu.com\r\n\r\n"
	conn.Write([]byte(httpRequest))
	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	if n == 0 {
		return
	}
	fmt.Printf("|%s|\n", string(buf[:n]))
}
