package main

import "fmt"

//递归调用
func test(n int) {
	if n > 2 {
		n--
		test(n)
	} else {
		fmt.Println("n=", n)
	}

}
func main() {
	test(4)
}
