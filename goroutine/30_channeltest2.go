package main

import "fmt"

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool)
	go writeData(intChan)
	go readData(intChan, exitChan)
	<-exitChan
	/*for{
		_,ok:=<-exitChan
		if !ok{
			break
		}
	}*/

}
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("write data:", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readdata 读到数据了=%v\n", v)
	}
	exitChan <- true
	close(exitChan)

}
