package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOROOT())
	n := runtime.GOMAXPROCS(4) //设置当前进程使用的最大cpu核数，返回上一次成功调用的设置值，首次调用返回默认值。比如：将cpu设置为1，返回是前一次的cpu的个数4
	fmt.Println("n=", n)

	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
