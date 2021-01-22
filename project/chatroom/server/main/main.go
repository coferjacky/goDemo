package main

import (
	"fmt"
	"net"
)

//1 监听
//2 等待客户端的连接
//3 初始化工作

//处理和客户端的通讯

/*func readPkg(conn net.Conn) (mes message.Message, err error) {
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
}*/

/*//处理登录请求逻辑
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
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

}*/

/*func writePkg(conn net.Conn, data []byte) (err error) {
	//发送一个长度給对方
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

	//发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn write(byte) fail", err)
	}
	return

}*/

//处理注册逻辑请求

//根据客户端发送信息种类不同，调用不同函数
/*func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {

	return
}*/

//获取套接字
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	//这里调用总控制，创建一个实例
	//循环的客户端发送的信息
	Processor := &Processor{
		Conn: conn,
	}
	err := Processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯的协程错误err", err)
		return
	}
	/*for {
	//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message,Err
	mes, err := readPkg(conn)
	//fmt.Println("mes=", mes)
	if err != nil {
		if err == io.EOF {
			fmt.Println("客户端退出，服务端也退出")
			return
		} else {
			fmt.Println("readpkg err=", err)
			return
		}

	}
	//fmt.Println("mes=", mes)
	err = serverProcessMes(conn, &mes)
	if err != nil {
		return
	}*/

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
