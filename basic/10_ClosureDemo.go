package main

import "fmt"

func main() {
	//由于没有用到闭包，所以这里的sum没有实现累加计数，导致每次sum调用均被初始化了

	for i := 0; i < 5; i++ {
		fmt.Printf("i=%d \t", i)
		fmt.Println(add2(i))
	}
}

func add2(i int) int {
	sum := 0
	sum += i
	return sum
}
