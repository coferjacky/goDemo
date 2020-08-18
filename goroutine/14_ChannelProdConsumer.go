package main

import (
	"fmt"
	"time"
)

/*
生产者与消费者模式实现
*/

func main() {
	ch := make(chan int, 5)
	go producer(ch)
	consumer(ch)
}
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("生产:", i)
		out <- i * i
	}
	close(out)
}
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("消费者拿到", num)
		time.Sleep(time.Second)
	}
}
