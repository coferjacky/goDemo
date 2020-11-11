package main

import "fmt"

//求1-200000的数字中，哪些是素数
/*
使用并发/并行方式，将统计素数的任务分配给多个（4个）goroutine去完成，完成任务时间短
注意：用map来保存primeNum

*/
func main() {

	intchan := make(chan int, 1000)
	primechan := make(chan int, 2000) //放结果
	exitchan := make(chan bool, 4)    //4个协程公用
	go putNum(intchan)
	for i := 0; i < 4; i++ {
		go primeNum(intchan, primechan, exitchan)
	}

	//主线程进行处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitchan
		}
		//管道取出结果4个，放心关闭 primeNum
		close(primechan)
	}()
	//遍历我们的primechan
	for {
		res, ok := <-primechan
		if !ok {
			break
		}
		fmt.Printf("素数有=%d\n", res)
	}

}

func putNum(intchan chan int) {
	for i := 0; i < 8000; i++ {
		intchan <- i
	}
	close(intchan)
}
func primeNum(intchan chan int, primechan chan int, exitchan chan bool) {
	//var num int
	var flag bool

	for {
		num, ok := <-intchan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primechan <- num
		}
	}
	fmt.Println("有个primeNum协程因为取不到数据，退出")
	exitchan <- true

}
