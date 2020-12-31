package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"project/chatroom/common/message"
)

//处理和客户端的通讯

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取到了客户端发送的数据 ...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	//根据读到的长度buf[:4]转换为一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint16(buf[0:4])
	//从conn里面读pkglen的字节放进buf里面
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail err=", err)
		return
	}
	//把pkg反序列化为message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err", err)
		return
	}
	return
}

//获取套接字
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	//循环的客户端发送的信息

	for {
		readPkg()
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
