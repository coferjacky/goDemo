package main

import "fmt"

/*
放A-Z到数组里面,然后检索出来
*/
func main() {

	var mychars [26]byte
	for i := 0; i < 26; i++ {
		mychars[i] = byte('A' + i)

	}
	for i, v := range mychars {
		fmt.Printf("下标是%d,值为%c\n", i, v)
	}
}
