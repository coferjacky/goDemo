/*
主go程结束，子go程随之结束，所以此处只执行4次
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("----i'm goroutine------")
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("--------i'm main--------")
		time.Sleep(time.Second) //定时时钟到达，唤醒线程，但是还是要与其他go程来争夺时间片
		if i == 2 {
			break
		}
	}
}
