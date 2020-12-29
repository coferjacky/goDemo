package main

import (
	"fmt"
	"net"
	"project/chatroom/common/message"
)

func login(userID int, userPwd string) (err error) {
	//1.链接服务器
	conn, err := net.Dial("tcp", "192.168.0.20:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//2.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginResMessType
	//3.创建一个loginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userID
	loginMes.UserName = userPwd
	//4 将loginmes进行序列化
	fmt.Println(loginMes)
}
