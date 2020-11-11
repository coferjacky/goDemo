package main

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
}

func putNum(intchan chan int) {
	for i := 0; i < 8000; i++ {
		intchan <- i
	}
	close(intchan)
}
func primeNum(intchan chan int, primechan chan int, exitchan chan bool) {
	var num int
	var flag bool
	for {
		num = <-intchan

		for i := 2; i < num; i++ {
			if num%i == 0 {

			}
		}
	}
}
