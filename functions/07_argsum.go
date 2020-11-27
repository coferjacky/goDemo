package main

import "fmt"

func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]

	}
	return sum
}
func main() {
	res4 := sum(10, 11, 12, 12)
	fmt.Println(res4)
}
