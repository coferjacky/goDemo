package main

/*
使用指针变量获取命令行的输入信息
*/
import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "", "procemode") //参数1：参数名称；参数2：参数默认值；参数说明：使用help时会出现再说明中
func main() {
	//解析命令行参数
	flag.Parse()

	fmt.Println(*mode)
}
