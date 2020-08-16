package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		defer fmt.Println("aaaaaaaaaaaaaaaaaa") //defer是在本函数结束时再调用,无论是return还是goexit
		test()
		test2()
		fmt.Println("bbbbbbbbbbbbbbbbbbbbb")
	}()
	for {

	}

}
func test() {
	defer fmt.Println("cccccccccccccccccc")
	return //返回当前函数调用，返回给调用者，后续指令不会执行了

	fmt.Println("DDDDDDDDDDDDDDD")
}
func test2() {
	defer fmt.Println("eeeeeeeeeeee")
	runtime.Goexit() //结束调用该函数的当前go程，goEXIT之前注册defer都会生效
	fmt.Println("ffffffffff")
}
