package main

import (
	"fmt"
)

//Monkey结构体
type Monkey struct {
	Name string
}

//LittleMonkey结构体
type LittleMonkey struct {
	Monkey //继承
}

//声明接口,虽然接口即便不声明，本文件代码也能运行正确。当时如果写了接口以后，相当于规定了扩展功能的设计。
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, " 生来会爬树..")
}

//让LittleMonkey实现BirdAble
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, " 通过学习，会飞翔...")
}

//让LittleMonkey实现FishAble
func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, " 通过学习，会游泳..")
}

func main() {

	//创建一个LittleMonkey 实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()

}
