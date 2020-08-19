package main

// sleep   newTimeer   After ---- time包下面的
/*
定时器的重置和停止
*/
import (
	"fmt"
	"time"
)

func main() {
	myTime := time.NewTimer(time.Second * 20) //创建定时器
	myTime.Reset(time.Second * 1)             //定时器的重置
	go func() {                               //

		<-myTime.C
		fmt.Println("子go程序读取定时完毕")
	}()
	//myTime.Stop() //设置定时器停止（定时器归零），<-myTime.C会发生阻塞。这里系统停止对channel进行写入，但是系统保持着写端，go程也仍然保持着读端，所以此时不死锁
	for {

	}

}
