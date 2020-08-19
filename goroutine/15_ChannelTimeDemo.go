package main

/*
	定时器 结构体如下
		type Time struct{
			C <-chan Time
   		 	r runtimeTime
		}

*/
import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器
	//方法1：sleep

	time.Sleep(time.Second)
	fmt.Println("当前时间", time.Now())
	//方法2：Timer.C
	myTimer := time.NewTimer(time.Second * 2) //创建定时器，指定定时时长
	nowTime := <-myTimer.C
	fmt.Println("当下时间：", nowTime)

	//方法3：time.After
	nowTime2 := <-time.After(time.Second * 2) //定时满系统自动写入系统时间
	fmt.Println("现在时间：", nowTime2)

}
