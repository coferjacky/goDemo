/*
注意：1 监听case中，没有满足监听条件的情况 阻塞
     2 监听的case中，有多个满足监听条件，任选一个执行
     3 可以使用default来处理所有case都不满足条件状况。通常这种情况不大用
     4 select 自身不带有循环机制，需要借助外层for来循环监听
     5  break只跳出select ,类似switch中的用法

*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)    //用来进行数据通信的channel
	quit := make(chan bool) //用来判断是否退出 channel
	go func() {             //写数据
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true //通知主go程退出
		runtime.Gosched()
	}()
	for { //主go程 读数据
		select {
		case num := <-ch:
			fmt.Println("读到：", num)

		case <-quit:
			//break		//break 跳出select
			//runtime.Goexit() 只能退子协程
			return //终止进程

		}
		fmt.Println("------------")
	}
}
