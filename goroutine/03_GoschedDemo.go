package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for {
			fmt.Println("this is goroutine test")

		}
	}()
	for {
		runtime.Gosched() //出让当前go程所占用的cpu时间片，当再次获得cpu以后，从出让位置继续恢复执行
		fmt.Println("this is main test")

	}
}
