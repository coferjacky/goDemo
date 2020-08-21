package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
读写锁的示例  读时共享，写时独占，写的优先级比读要高  ，下例是读写锁和channel混用导致的死锁，运行报错
注意：尽量不要将互斥锁、读写锁与channel混用，会发生隐形死锁
*/
var rmMutex sync.RWMutex //s锁只有一把，2个属性r 和 w
func main() {

	//播随机数种子
	rand.Seed(time.Now().UnixNano())
	//quit:=make(chan bool)  //用于关闭主go进程的channel
	quit := make(chan bool)
	ch := make(chan int) //用于数据传递的channel
	for i := 0; i < 5; i++ {

		go readGo(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go writeGo(ch, i+1, quit)
	}
	<-quit //这里不用for循环 优雅很多
}
func readGo(in <-chan int, idx int) {
	rmMutex.RLock() //对共享数据 channel加锁
	num := <-in
	fmt.Printf("<-----%d次读go程，读出%d\n", idx, num)
	rmMutex.RUnlock() //对共享数据channel解锁
}
func writeGo(out chan<- int, idx int, quit chan<- bool) {
	//生成随机数

	num := rand.Intn(1000)
	rmMutex.Lock() //写模式加锁
	out <- num
	fmt.Printf("---->%d次写go程，写入%d\n", idx, num)
	time.Sleep(time.Millisecond) //放大实验现象
	rmMutex.Unlock()             //写模式解锁
	quit <- true

}
