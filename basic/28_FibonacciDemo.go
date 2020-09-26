package main

import "fmt"

//Fibonacci
func main() {
	fmt.Print(fib(7))
}
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
