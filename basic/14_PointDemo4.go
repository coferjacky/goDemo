/*
指针中的指针
*/
package main

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int
	a = 1234
	ptr = &a
	fmt.Println("ptr", ptr)

	pptr = &ptr
	fmt.Println("pptr", pptr)
	fmt.Printf("变量 a=%d \n", a)
	fmt.Printf("指针变量*ptr=%d \n", *ptr)
	fmt.Printf("指向指针的制作变量 **pptr=%d \n", **pptr)
}
