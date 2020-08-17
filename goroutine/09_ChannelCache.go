package main

/*
有缓冲channel 是-----异步通信   len是channel中的剩余的元素，cap是channel中元素容量
*/
import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3) //存满3个元素前不会阻塞
	fmt.Println("len=", len(ch), "cap=", cap(ch))
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
			fmt.Println("子go程中：i", i)
		}
	}()

	time.Sleep(time.Second * 3)

	for i := 0; i < 8; i++ {
		num := <-ch
		fmt.Println("主go程读到：", num)
	}

}
