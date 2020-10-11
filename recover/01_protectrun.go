package main

//宕机恢复的一枚例子
import (
	"fmt"
	"runtime"
)

//崩溃时需要传递的上下文信息
type panic_Context struct {
	function string
}

//当panic触发崩溃时，该函数将结束运行，此时defer后的必报
func Protect_Run(entry func()) {
	defer func() {
		//发生宕机时，获取panic传递的上下文并打印
		err := recover()

		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime err:", err) //运行时错误

		default:
			fmt.Println("error:", err) //非运行时错误

		}
	}()
	entry()
}
func main() {
	fmt.Println("运行前")

	//手工触发错误
	Protect_Run(func() {
		fmt.Println("手工宕机前")
		//使用panic传递上下文
		panic(&panic_Context{"手工触发panic"})
		//下面的不执行了
		fmt.Println("手工宕机后")
	})
	Protect_Run(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1

		fmt.Println("赋值宕机后")
	})
	fmt.Println("运行后")

}

/*
1 有panic没有recover，程序宕机
2 有panic 也有recover 程序不会宕机，执行完对应的defer后，从宕机点退出当前函数后继续执行
*/
