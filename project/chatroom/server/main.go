package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"project/chatroom/common/message"
)

//处理和客户端的通讯

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取到了客户端发送的数据 ...")
	//conn.read在conn没有关闭的情况下才会阻塞，
	//如果客户端关闭conn则不会阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {

		//err=errors.New("read pkg header error")
		return
	}
	//根据读到的长度buf[:4]转换为一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//从conn里面读pkglen的字节放进buf里面
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err=errors.New("read pkg body error")
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
		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message,Err
		mes, err := readPkg(conn)

		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出")
				return
			} else {
				fmt.Println("readpkg err=", err)
				return
			}

		}
		fmt.Println("mes=", mes)
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
			return
		}
		//连接成功，则启动一个协程和客户端保持交互和通讯
		go process(conn)
	}
}
