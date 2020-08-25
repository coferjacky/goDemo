package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	list := os.Args
	//提起文件的绝对路径
	if len(list) != 2 {
		fmt.Println("格式为：go run xxxx.go文件绝对路径")
		return
	}
	//提取文件的绝对路径
	filePath := list[1]
	//提取文件名
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("filepath err:", err)
		return
	}
	fileName := fileInfo.Name()
	//主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8009")
	if err != nil {
		fmt.Println("conn err:", err)
		return
	}
	defer conn.Close()
	//发送文件名给 接受端
	conn.Write([]byte(fileName))
	//读取服务器回发的OK 状态
	buf := make([]byte, 16)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	if "ok" == string(buf[:n]) {
		//写文件内容给服务器---借助conn
		sendFile(conn, filePath)
	}
}

func sendFile(conn net.Conn, filePath string) {
	//只读打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	defer f.Close()
	//从文件中读数据，写给网络接受者，读多少 写多少
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {

			if err == io.EOF {
				fmt.Println("发送文件完成")
			} else {
				fmt.Println("read err:", err)

			}
			return
		}
		//写到网络socket中:从本地读 从网络写
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn err:", err)
			return
		}
	}
}
