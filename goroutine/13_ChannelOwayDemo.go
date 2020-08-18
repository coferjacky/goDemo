package main

import "fmt"

/*
单向d写和读
     channel  var sendCh chan<- int   ||  sendCh=make(chan<- int)  单向写不能进行读操作
     channel var recvCh <-chan int    ||  recvCh=make(<-chan int)  单向读不能进行写操作
      双向chanel可以赋值给单向，单向不能赋值给双向
      传参时是传引用

*/
func main() {
	ch := make(chan int)
	go func() {
		send(ch)

	}()
	recv(ch)
}
func send(out chan<- int) {
	out <- 8899
	close(out)
}
func recv(in <-chan int) {
	n := <-in
	fmt.Println("读到", n)
}
