package main

//模拟echo 打印每个参数和索引 每个一行
import (
	"fmt"
	"os"
)

func main() {

	for index1, arg := range os.Args[1:] {

		fmt.Println(index1, arg)
	}

}
