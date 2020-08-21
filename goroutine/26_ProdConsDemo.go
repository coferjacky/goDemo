package main

/*
条件变量的使用：1 加锁  2 访问公共区  3 解锁 4 唤醒阻塞在条件变量的对端
  使用条件变量cond ，首先确定用什么类型锁，比如用的是 互斥锁，先做lock操作，判断 满足条件就做wait操作
  wait函数做三件事：1 阻塞等待条件变量满足  2 释放已掌握的互斥锁相当于cond.L.Unlock 原子性
                   3 当被唤醒,wait()函数返回，解除阻塞并重新获得互斥锁 相当于cond.L.Lock
  signal 函数来唤醒对端（对端一个GO程） 阻塞的go程
  BroadCast  函数广播唤醒，产生惊群效应

*/
import "fmt"

func main() {
	product := make(chan int)

	go producer1(product)
	consumer1(product)

	for {

	}
}
func producer1(out chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("生产者生产：%d\n", i)
		out <- i
	}
	close(out)

}
func consumer1(in <-chan int) {

	for num := range in {
		fmt.Printf("消费者消费数据：%d\n", num)
	}
}
