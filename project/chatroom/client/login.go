package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"project/chatroom/common/message"
	"time"
)

func login(userID int, userPwd string) (err error) {
	//1.链接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	//2.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginResMessType
	//3.创建一个loginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userID
	loginMes.UserPwd = userPwd
	//4 将loginMes结构体进行序列化,返回loginMes的json编码 即[]byte
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5 把data给结构体  转换为string类型
	mes.Data = string(data)
	//6 将mes进行序列化(相当于返回)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marchal err=", err)
		return
	}
	//7 到这个时候data在进行发送
	//7.1先把data的长度发送给服务器
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write ERR", err)
		return
	}
	fmt.Printf("客户端发送消息的长度成功 发送长度是%d,内容是%s", len(data), string(data))
	//  return

	//发送消息
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.write(data) fail", err)
		return
	}
	//休眠20s
	time.Sleep(20 * time.Second)
	fmt.Println("休眠20...")
	//这里还需要处理服务器端返回的消息
	return

}
