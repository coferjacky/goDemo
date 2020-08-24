package main

//并发服务器
import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//第一步：创建监听套接字
	listener, err := net.Listen("tcp", "192.168.0.101:8001")
	if err != nil {
		fmt.Println("net.Lister err:", err)
		return
	}
	//第二步：defer close
	defer listener.Close()

	//第三步：监听客户端连接请求
	for {
		fmt.Println("服务器等待客户端连接。。。。")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept:", err)
			return
		}

		//第四步：创建go程 具体完成服务器客户端的数据通信
		go HandlerConnect(conn)
	}
}

//第五步：实现handlerConnect
func HandlerConnect(conn net.Conn) {

	defer conn.Close()
	//获取连接客户端的address
	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端成功连接")
	//循环读取客户端发送数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)

		if "exit\n" == string(buf[:n]) || "exit\r\n" == string(buf[:n]) {
			fmt.Println("服务器接收到客户端请求，服务器关闭")
			return
		}
		if n == 0 {
			fmt.Println("服务器检测到客户端已经关闭，服务器断开连接")
			return
		}
		if err != nil {
			fmt.Println("conn.read err:", err)
			return
		}
		fmt.Println("服务器读到数据：", string(buf[:n])) //服务器使用数据
		//小写转大写回发客户端
		//strings.ToUpper(string(buf[:n])) 这是把小写转大写
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}
