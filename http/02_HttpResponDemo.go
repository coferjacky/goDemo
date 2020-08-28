package main

import "net/http"

func handle(w http.ResponseWriter, r *http.Request) {

}
func main() {
	//第一步 注册回调函数，该回调函数会在服务器被访问时，自动被调用
	http.HandleFunc("/itcast", handle)
	//第二步 绑定服务器监听的地址
	http.ListenAndServe("127.0.0.1:8000", nil) //第二个参数我们通常传nil
}
