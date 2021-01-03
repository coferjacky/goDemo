package process

import (
	"encoding/json"
	"fmt"
	"net"
	"project/chatroom/common/message"
)

type UserProcess struct {
	//字段
	Conn net.Conn
}

//处理登录请求逻辑
func (this *UserProcess) serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//1 从mes中取出mes.Data,并直接反序列号为LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	//申明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//再声明一个LoginResMes
	var loginResMes message.LoginResMes

	//如果用户的id为100，密码是123456 ，认为是合法的，否则报错
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200

	} else {
		//500表示该用户不存在
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在 请住持再使用"
	}

	//3 将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.marshar fail", err)
		return
	}
	//4 将data赋值給resMes
	resMes.Data = string(data)

	//5  resMes进行序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshar fail", err)
		return
	}
	//6 发送data 将他封装到writePkg函数里面
	err = writePkg(conn, data)
	return

}
