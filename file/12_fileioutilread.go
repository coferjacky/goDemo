package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "c:/myjack.abc"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	//把读取到内容显示到终端
	fmt.Printf("%v", string(content))
}
