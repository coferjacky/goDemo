package main

import (
	"fmt"
)

/*
channel关闭后 依然可以进行读，但是不能写了
*/
func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
		}
		close(ch)
	}()
	for {
		if num, ok := <-ch; ok == true {
			fmt.Println("读到数据了:", num)
		} else {
			n := <-ch
			fmt.Println("关闭后：读到数据：", n) //管道有个特性，只要你读到0则 说明管道对端（写端）已经关闭了
			break
		}
	}

}
