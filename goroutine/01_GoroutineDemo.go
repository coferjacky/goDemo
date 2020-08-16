package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 50; i++ {
		fmt.Println("-----唱歌中-------")
		//time.Sleep(100 * time.Millisecond)
	}
}
func dance() {
	for i := 0; i < 50; i++ {
		fmt.Println("-----跳舞中-------")
		//time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	go sing()
	go dance()
	for {
		fmt.Println("---主线程来了-----")
		time.Sleep(100 * time.Millisecond)
	}
}
