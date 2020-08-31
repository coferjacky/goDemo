package main

import (
	"fmt"
	"net/http"
)

func myHandle(w http.ResponseWriter, r *http.Request) {
	//w:写给客户端的数据内容
	w.Write([]byte("this is out web service"))

	//r:从客户端读到的内容
	fmt.Println("head:", r.Header)
	fmt.Println("url", r.URL)
	fmt.Println("Methed:", r.Method)
	fmt.Println("Host:", r.Host)
	fmt.Println("remoteAdd:", r.RemoteAddr)
	fmt.Println("body:", r.Body)
}
func main() {
	//注册回调函数，函数再客户端访问服务器时，自动调用
	http.HandleFunc("/itcast", myHandle)
	//绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}
