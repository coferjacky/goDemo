package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//创建用户结构体类型
type Client struct {
	C    chan string
	Name string
	Addr string
}

//创建全局map 存储在线用户,先声明一个map变量
var onlineMap map[string]Client

//创建全局channel 用来传递用户的消息message
var message = make(chan string)

func main() {
	//创建监听套接字
	listener, err := net.Listen("tcp", "192.168.0.80:8000")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()
	//创建管理者go程 管理map和message
	go Manageer()

	//监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			return
		}
		//启动go程处理客户端数据请求
		go HandlerConnect(conn)
	}
}

func Manageer() {
	//初始化onlineMap
	onlineMap = make(map[string]Client)

	//监听全局channel中是否有数据,有数据存在msg，无数据堵塞
	for {
		msg := <-message

		//循环发送消息给所有在线用户,要想执行，必须msg:=<message执行完毕，解除阻塞
		for _, clnt := range onlineMap {
			clnt.C <- msg
		}
	}
}
func MakeMsg(clnt Client, msg string) (buf string) {
	buf = "[" + clnt.Addr + "]" + clnt.Name + ":" + msg
	return
}

func HandlerConnect(conn net.Conn) {

	defer conn.Close()
	//创建channel 判断用户是否活跃

	hasData := make(chan bool)

	//获取用户的网络地址和端口号
	netAddr := conn.RemoteAddr().String()

	//创建新连接用户的结构体信息,默认用户名是ip+addr
	clnt := Client{make(chan string), netAddr, netAddr}
	//将新链接的用户添加的map中,key：IP+Port,value:client
	onlineMap[netAddr] = clnt

	//创建专门用来给当前用户发送消息的go程
	go WriteMsgClient(clnt, conn)

	//发送用户上线的消息给massage通道中
	//message<-"["+netAddr+"]"+clnt.Name+"  login   "
	message <- MakeMsg(clnt, "login")

	//创建一个channel 判断用户退出状态
	isQuit := make(chan bool)

	//创建一个匿名go程，专门处理用户发送信息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println("检测到客户端：%s 退出\n", clnt.Name)
			}
			if err != nil {
				fmt.Println("conn err", err)
				return
			}
			//将读入的用户消息保存到msg中，string类型
			msg := string(buf[:n-1]) //带了个\n则所以要减去一个1
			//提取在线用户列表
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("user List:\n"))
				//遍历当前map 获取在线用户
				for _, user := range onlineMap {
					userInfo := user.Addr + ":" + user.Name + "\n"
					conn.Write([]byte(userInfo))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" { //reanme |
				newName := strings.Split(msg, "|")[1] //msg[8:]
				clnt.Name = newName                   //修改结构体成员name
				onlineMap[netAddr] = clnt             //更新在线用户列表
				conn.Write([]byte("rename successful\n"))
			} else {

				//将读到用户消息，写入到message
				message <- MakeMsg(clnt, msg)
			}

			hasData <- true

		}
	}()

	//确保这个进程不退出
	for {
		//监听channel上的数据流动
		select {
		case <-isQuit: //将用户从在online移除
			delete(onlineMap, clnt.Addr)
			message <- MakeMsg(clnt, "logout") //写入用户退出消息到全局channel
			return
		case <-hasData:
			//什么都不做，目的是重置下面的计时器
		case <-time.After(time.Second * 10): //将用户从在online移除
			delete(onlineMap, clnt.Addr)
			message <- MakeMsg(clnt, "time out for logout") //写入用户退出消息到全局channel
			return
		}

	}
}

func WriteMsgClient(clnt Client, conn net.Conn) {
	//监听用户自带的channel上面是否有消息
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}

}
