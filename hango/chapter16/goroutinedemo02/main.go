package main

import (
	"fmt"
	"time"
)

var (
	myMap = make(map[int]int, 10)
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res
}

func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 10)
	for k, v := range myMap {
		fmt.Printf("map[%d]=%d\n", k, v)
	}
}
