package main

//可以看到调用不影响指针赋值
import "fmt"

func main() {
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}

func incr(p *int) int {
	*p++
	return *p
}
