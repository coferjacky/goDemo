package main

/*
同一时间只有一个人共享数据
*/
import (
	"fmt"
	"sync"
	"time"
)

/*  使用channel的方式进行打印
var ch=make(chan int)
func main() {
	go person11()
	go person21()
	for{
		;
	}
}
func printer1(str string){
	for _,ch:=range str{
		fmt.Printf("%c",ch)
		time.Sleep(time.Microsecond*300)
	}
}
func person11()  {
	printer1("hello")
	ch<-1
}
func person21()  {
	<-ch
	printer1("world")

}*/

var mutex sync.Mutex //创建一个互斥量，新建的互斥锁状态位0 未加锁。锁只有一把，go程与go程之间锁只有一把

func main() {
	go person11()
	go person21()
	for {

	}
}
func printer1(str string) {
	mutex.Lock() //共享数据加锁
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Microsecond * 300)
	}
	mutex.Unlock() //共享数据释放锁
}
func person11() {
	printer1("hello")

}
func person21() {

	printer1("world")

}
