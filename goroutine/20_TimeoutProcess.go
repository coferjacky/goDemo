package main

/*
延时处理
*/
import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num:", num)
			case <-time.After(3 * time.Second):
				quit <- true
				//break   只从select跳出来
				goto lable
			}
		}
	lable:
		fmt.Println("break to lable------")
	}()
	for i := 0; i < 2; i++ {
		ch <- i
		time.Sleep(time.Second * 2)

	}
	<-quit
	fmt.Println("finished!!!")
}
