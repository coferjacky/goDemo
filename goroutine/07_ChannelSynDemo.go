package main

/*

channel同步-传递数据
*/
import "fmt"

func main() {

	ch := make(chan string) //无缓存channel
	//ch<-"子go打印完毕"   //放这里就死锁了
	// len： channel中剩余未读数据个数，cap ：channel中的容量
	fmt.Println("len(ch)=", len(ch), "cap(ch)=", cap(ch))

	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("i=", i)
		}
		//通知主go打印完毕
		ch <- "子go打印完毕"
	}()
	str := <-ch
	fmt.Println("str=", str)
}
