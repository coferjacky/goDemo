package main

import (
	"encap/model"
	"fmt"
)

func main() {
	accout := model.NewAccount("caofei", "555555", 40)
	if accout != nil {

		fmt.Println("创建成功了", accout)
	} else {
		fmt.Println("失败了")
	}

	accout.SetPwd("777777")

	pwd := accout.GetPwd()
	fmt.Println("pwd=", pwd)
}
