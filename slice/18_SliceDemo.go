package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	arr1 := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	s := arr[1:3:5] //low:high:max low第角标，最高角标，容量=max-low容量
	s1 := arr1[3:5] //没有指定max 则用原数组的长度来算
	fmt.Println("s=", s)
	fmt.Println("len(s)=", len(s))
	fmt.Print("cap(s)=", cap(s))

	fmt.Println("s=", s1)
	fmt.Println("len(s)=", len(s1))
	fmt.Print("cap(s)=", cap(s1))

}
