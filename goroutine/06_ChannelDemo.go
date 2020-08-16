package main

import (
	"fmt"
	"time"
)

/*
   make(chan 在channel中传递的数据类型 容量)   容量=0： 无缓冲channel(可省)
   每当有一个进程启动时，系统会自动打开三个文件：标准输入、标准输出、标准错误 ---- 对应三个文件
   stdin
   stdout
   stderr
*/

//全局定义channel,用来完成数据同步
var channel = make(chan int)

//定义一个打印机
func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}
}

//定义两个人
func person1() { //person1先执行
	printer("hello")
	channel <- 8 //执行到这里时 由于传入channel的数据没有取出，所以进程被阻塞
}
func person2() { //person2后执行
	<-channel //执行到这里 被阻塞,person1 可能还没有执行到 channel写入位置
	printer("world")
}

func main() {
	go person1()
	go person2()
	for {

	}
}
