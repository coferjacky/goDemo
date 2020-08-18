package main

import (
	"fmt"
	"time"
)

/*
读有缓存channel：如果缓存区内有数据，则先读数据，读完数据后可以继续读，直到读到0
*/
func main() {
	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	time.Sleep(time.Second * 10)
	fmt.Println("子GO结束了")
	/*for{
		if num,ok:=<-ch;ok==true{
			fmt.Println("读到数据：",num)
		}else {
			n:=<-ch
			fmt.Println("关闭后还能读到数据：",n)
			break
		}
	}*/
	fmt.Println("------采用range也可以-------")
	for num := range ch {
		fmt.Println("读到数据：", num)
	}
}
