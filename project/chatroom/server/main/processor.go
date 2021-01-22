package main

import (
	"fmt"
	"io"
	"net"
	"project/chatroom/common/message"
	"project/chatroom/server/process"
	"project/chatroom/server/utils"
)

//创建一个Processor的结构体
type Processor struct {
	Conn net.Conn
}

//根据客户端发送信息种类不同，调用不同函数
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//根据message处理一个登录后信息返回
		//创建一个userProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}

		err = up.ServerProcessLogin(this.Conn, mes)
	case message.RegisterMesType:
		//处理注册逻辑
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

//获取套接字
func (this *Processor) process2() (err error) {

	//循环的客户端发送的信息

	for {

		//创建Transfer实例完成读包的任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message,Err
		mes, err := tf.ReadPkg()
		//fmt.Println("mes=", mes)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出")
				return err
			} else {
				fmt.Println("readpkg err=", err)
				return err
			}

		}
		//fmt.Println("mes=", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}

	}

}
