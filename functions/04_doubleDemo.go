package main

import "fmt"

//只要首先命名double的返回值，再增加defer语句，我们就可以再每次double被调用时 输出参数以及返回值

func main() {
	_ = double(4) //
}
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d)=%d", x, result) }()
	return x + x
}
