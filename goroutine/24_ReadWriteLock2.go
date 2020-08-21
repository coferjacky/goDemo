package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
读写锁的示例  读时共享，写时独占，写的优先级比读要高  ，下例是读写锁的运用
注意：尽量不要将互斥锁、读写锁与channel混用，会发生隐形死锁
*/
var rmMutex1 sync.RWMutex //s锁只有一把，2个属性r 和 w
var value int             //定义全局变量，模拟共享数据

func main() {

	//播随机数种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {

		go readGo1(i + 1)
	}
	for i := 0; i < 5; i++ {
		go writeGo1(i + 1)
	}
	for {

	}
}
func readGo1(idx int) {
	for {

		rmMutex1.RLock() //对共享数据 channel加锁
		num := value     //写入全局value变量里面
		fmt.Printf("<-----%d次读go程，读出%d\n", idx, num)
		rmMutex1.RUnlock() //对共享数据channel解锁
	}
}
func writeGo1(idx int) {

	for {

		//生成随机数
		num := rand.Intn(1000)
		rmMutex1.Lock() //写模式加锁
		value = num
		fmt.Printf("---->%d次写go程，写入%d\n", idx, num)
		//time.Sleep(time.Millisecond)  //放大实验现象
		rmMutex1.Unlock() //写模式解锁
	}

}
