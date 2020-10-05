package main

import "fmt"

func main() {
	fmt.Println(equal(map[string]int{"a": 0}, map[string]int{"a": 1}))
}

//判断两个map是否相等是不行的，但是可以判断两个map是否包含相同的key和value
func equal(x, y map[string]int) bool {
	if len(x) != len(y) { //如果长度不等返回false
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}

	}
	return true
}
