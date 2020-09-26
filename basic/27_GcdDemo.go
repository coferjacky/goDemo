package main

//求两数的最大公约数
import "fmt"

func main() {
	fmt.Println(gcd(12, 18))
}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
