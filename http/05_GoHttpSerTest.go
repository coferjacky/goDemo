package main

import (
	"fmt"

	"net/http"
	"os"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("客户端请求", r.URL)
	OpenSendFile(r.URL.String(), w)
}

func OpenSendFile(fName string, w http.ResponseWriter) {
	pathFileName := "C:/" + fName
	f, err := os.Open(pathFileName)
	if err != nil {
		w.Write([]byte("no such file or dir"))
		fmt.Println("OPERR", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 4096)
	for {
		n, _ := f.Read(buf) //从本地将文件内容读取
		if n == 0 {
			return
		}
		w.Write(buf[:n])

	}
}

func main() {
	//注册回调函数
	http.HandleFunc("/", myHandler)
	//绑定监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}
