package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("helloworld")
	}
}
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "GOLANG" //ERROR map必须先make再用
}
func main() {
	go sayHello()
	go test()
	for {
		time.Sleep(time.Second)
		fmt.Println("hello")
	}
}
