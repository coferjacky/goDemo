package main

import (
	"fmt"
)

type Person3 struct {
	name     string
	age      int
	flg      bool
	interset []string
}

//初始化结构体方式1：通过函数的参数初始化结构体
func initFunction(p *Person3) {
	p.name = "nami"
	p.age = 67
	p.flg = true
	p.interset = append(p.interset, "shopping")
	p.interset = append(p.interset, "playingtvgame")
}

//初始化结构体方式2：通过函数返回值来初始化结构体
func initFunction2() *Person3 {
	p := new(Person3)
	p.name = "jackbaul"
	p.age = 18
	p.flg = true
	p.interset = append(p.interset, "shopping1")
	p.interset = append(p.interset, "playingtvgame1")
	return p // 注意：不能返回局部变量的地址值，因为函数调用结束后会丢失。所以只能返回局部变量的值
}
func main() {
	var person Person3

	initFunction(&person) //该函数给person 做初始化
	fmt.Println("person:", person)
	p2 := initFunction2()
	fmt.Println("person2:", *p2)
}
