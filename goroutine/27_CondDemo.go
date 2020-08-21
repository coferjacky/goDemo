package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
步骤
1 先创建条件变量var cond sync.Cond
2 指定条件变量配套的锁 cond.L=NEW(sync.Mutex)
3 给公共区加锁  （互斥锁）
4 判断是否到达阻塞条件（缓冲区满/空）  ---for循环  for len(ch)==cap(ch) {cond.wait()}
5 访问公共区--读 写数据 打印
6 解锁条件变量用的锁  cond.L.Unlock()
7 唤醒阻塞在条件变量上的对端

*/
var cond sync.Cond //创建全局条件变量

func main() {
	rand.Seed(time.Now().UnixNano()) //随机设置种子
	//设置5个消费者
	//quit:=make(chan bool)  //创建用于结束的通信的channel
	product := make(chan int, 3) //产品区采用channel模拟

	cond.L = new(sync.Mutex) //指定条件变量和对应的锁

	for i := 0; i < 5; i++ {
		go consumer3(product, i+1)
	}
	//设置3个生产者
	for i := 0; i < 3; i++ {
		go producer3(product, i+1)
	}
	for {

	}
}
func producer3(out chan<- int, idx int) {
	for {
		cond.L.Lock()       //条件变量对应互斥锁加锁
		for len(out) == 3 { //产品区满 等待消费者消费
			cond.Wait() //把自己挂起，等条件变量满足，被消费者唤醒
		}
		num := rand.Intn(1000) //产生一个随机数
		out <- num             //开始生成，随机数进channel
		fmt.Printf("%d号生成者，产生数据%3d,公共区剩余%d 个数据\n", idx, num, len(out))
		cond.L.Unlock()         //生成结束，解锁互斥锁
		cond.Signal()           //唤醒阻塞的消费者
		time.Sleep(time.Second) //能让其他协程有执行机会
	}
}

func consumer3(in <-chan int, idx int) {
	for {
		cond.L.Lock()
		for len(in) == 0 {
			cond.Wait()
		}
		num := <-in //将生产区数据取走
		fmt.Printf("%d号消费者，消费数据%3d,公共区剩余%d 个数据\n", idx, num, len(in))
		cond.L.Unlock()
		cond.Signal() //唤醒阻塞的生产者
		time.Sleep(time.Millisecond * 500)
	}
}
