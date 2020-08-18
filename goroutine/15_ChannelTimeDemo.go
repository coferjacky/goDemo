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
	myTime := time.NewTimer(time.Second * 2)
	fmt.Println(myTime)
}
