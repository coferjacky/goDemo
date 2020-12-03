package main

import "fmt"

func main() {
	var chan1 chan int   //可读可写
	var chan2 chan<- int //只读
	chan1 = make(chan int, 3)
	chan2 = make(chan int, 3)
	chan1 <- 2
	chan2 <- 2
	//var chan3 <-chan
	//不能读 num:=<-chan2
	num := <-chan1
	fmt.Printf("chan2=", num)
}
