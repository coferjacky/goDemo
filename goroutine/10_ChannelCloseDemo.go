package main

import "fmt"

/*
使用close(ch)关闭channel
对端可以判断channel是否关闭
总结：数据不发送完毕，不应该关闭
已经关闭的channel不能再向其写数据
*/
func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
		}
		close(ch) //写端，写完数据主动关闭channel
		// ch<-666   panic: send on closed channel  已经关闭的channel不能再向其写数据
	}()
	for {
		if num, ok := <-ch; ok == true {
			fmt.Println("读到数据：", num)
		} else {
			break
		}
	}
	for {

	}
}
