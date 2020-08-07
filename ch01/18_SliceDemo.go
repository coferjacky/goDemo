package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}

	s := arr[1:3:5] //low:high:max low第角标，最高角标，容量=max-low容量
	fmt.Println("s=", s)
	fmt.Println("len(s)=", len(s))
	fmt.Print("cap(s)=", cap(s))
}
