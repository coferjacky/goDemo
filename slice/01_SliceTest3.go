package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := a[1:3] //2,3
	b = append(b, 5)
	fmt.Print(b) //2,3,5
	fmt.Print(a)
}
