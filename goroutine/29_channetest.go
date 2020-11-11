package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intchan值 %v\n", intChan)
	fmt.Printf("intChan地址:%v\n", &intChan)

	//定义一个可以存放任意数据类型的管道 3个数据
	var allChan chan interface{}
	allChan = make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tomjack"
	Cat4 := Cat{"小花猫", 4}
	allChan <- Cat4
	<-allChan
	<-allChan
	newCat := <-allChan

	fmt.Printf("newcat=%T,newCat=%v", newCat, newCat)

	//  fmt.Printf("newCat.name=%v",newCat.Name) 这种写法，编译器会将通道取出来的数据看成一个interfaced导致编译错误
	a := newCat.(Cat) //利用断言把类型转成猫猫
	fmt.Printf("newcatName=%v", a.Name)
}
