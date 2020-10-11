package main

//defer的练习
/*
每次bigSlowOperation()调用程序都会记录函数进入退出和持续时间
*/

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() //函数被调用时执行了trace的函数,然后sleep完毕后执行 return func中的内容

	time.Sleep(10 * time.Second)
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s(%s)", msg, time.Since(start))
	}
}
