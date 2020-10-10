package main

//当f（0）被调用时，发生panic异常，之前被延迟执行的3个fmt.printf被调用
import "fmt"

/*func main() {
	f(3)
}*/
func f(x int) {
	fmt.Println("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
