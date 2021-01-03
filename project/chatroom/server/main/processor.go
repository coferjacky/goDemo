package main

import (
	"fmt"
	"net"
	"project/chatroom/common/message"
)

//根据客户端发送信息种类不同，调用不同函数
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//根据message处理一个登录后信息返回
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
		//处理注册逻辑
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}
