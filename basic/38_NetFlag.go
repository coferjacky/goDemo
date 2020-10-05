package main

import (
	"fmt"
)

//练习对于iota的理解
/*
iota是初始是常量0，后面每行加1，使用这些常量来测试 设置和清理对应的bit位的值
*/
type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMultcast
)

func main() {
	var v Flags = FlagMultcast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v) //此时的v为 10000
	fmt.Printf("%b %t\n", v, IsUp(v))

}

func SetBroadcast(f2 *Flags) {
	*f2 |= FlagBroadcast //10000|=00010 ，v=10010
}

//相当于给第n位置0
func TurnDown(f2 *Flags) {
	*f2 &^= FlagUp //先*f2 10001  &^：FlagUp原来是00001，10001&^00001=10000 取值取决于FlagUp
}

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}
