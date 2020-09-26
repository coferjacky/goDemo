package main

//echo  当程序运行时，必须在使用标志参数对应变量之前先调用flag.Parse的函数，用于更新每个标志参数对应变量的值
import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newLine")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

}
