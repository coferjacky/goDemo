package main

import "fmt"

/*
请求出一个数组最大值，然后得出最大值对应下标

*/
func main() {
	var intArr = [...]int{1, 55, -2, 8, 11}
	maxVal := intArr[0]
	maxValIndex := 0
	for i := 0; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}

	}
	fmt.Printf("max最大值的index=%d,最大值为=%d", maxValIndex, maxVal)
}
