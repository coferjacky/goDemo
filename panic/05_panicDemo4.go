package main

import (
	"errors"
	"fmt"
)

func readconfig(str string) (err error) {
	if str == "ss.conf" {
		return nil
	} else {
		return errors.New("没有找到文件")
	}
}

func main() {
	err := readconfig("ss.conf")
	if err != nil {
		panic(err)
	}
	fmt.Println("正常执行")
}
