package main

import "fmt"

/*
生成一个斐波拉切数列放到切片中
*/
func fbn(n int) []uint64 {
	fbnslice := make([]uint64, n)
	fbnslice[0] = 1
	fbnslice[1] = 1
	for i := 2; i < n; i++ {
		fbnslice[i] = fbnslice[i-1] + fbnslice[i-2]
	}
	return fbnslice

}

func main() {

	fbnum := fbn(10)
	fmt.Println(fbnum)

}
