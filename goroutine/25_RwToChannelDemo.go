package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
用channel去替代读写锁，防止隐形锁
*/
func main() {
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go readGo2(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go writeGo2(ch, i+1)
	}
	for {

	}
}
func readGo2(in <-chan int, idx int) {
	for {
		num := <-in
		fmt.Printf("%dth 读GO程，读入：%d\n", idx, num)
	}
}
func writeGo2(out chan<- int, idx int) {
	for {
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("%dth 写GO程，写入：%d\n", idx, num)
		time.Sleep(time.Millisecond * 300)
	}
}
