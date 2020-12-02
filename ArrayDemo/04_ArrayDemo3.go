package main

//翻转打印数组
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intArr3 [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr3); i++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println(intArr3)
	for i := 0; i < len(intArr3)/2; i++ {
		//intArr3[i]第i元素，intArr3[len(intArr3)-1-i]倒数第i个元素
		intArr3[i], intArr3[len(intArr3)-1-i] = intArr3[len(intArr3)-1-i], intArr3[i]

	}
	fmt.Println(intArr3)
}
