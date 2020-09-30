package main

import (
	"fmt"
)

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
}

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}
