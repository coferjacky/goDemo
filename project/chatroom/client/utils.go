package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"project/chatroom/common/message"
)

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
