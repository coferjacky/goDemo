package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := &http.Client{}
	//创建一个http请求
	req, err := http.NewRequest("POST", "http://www.163.com/", strings.NewReader("key=value"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	//为标头添加信息
	req.Header.Add("User-Agent", "my Client")
	//开始请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	//读取服务器返回的内容
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	defer resp.Body.Close()

}
