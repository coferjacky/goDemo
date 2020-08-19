package main

/*
定时器周期运行
*/
import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)
	fmt.Println("now:", time.Now())
	timer := time.NewTicker(time.Second * 3) //定义一个定时器周期运行
	i := 0
	go func() {
		for {
			nowTime := <-timer.C
			i++
			fmt.Println("nowTime:", nowTime)
			if i == 8 {
				quit <- true
				break //return //runtime //goexit 效果一样
			}
		}
	}()
	<-quit
}
