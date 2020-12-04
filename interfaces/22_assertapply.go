package main

import "fmt"

type Usb1 interface {
	Start()
	Stop()
}

type Phone1 struct {
	Name string
}
type Camera1 struct {
	Name string
}

func (c Phone1) call() {
	fmt.Println("call来了")
}

func (c Phone1) Start() {
	fmt.Println("手机开始开始工作")
}
func (c Phone1) Stop() {
	fmt.Println("手机停止工作")
}
func (c Camera1) Start() {
	fmt.Println("照相机开始开始工作")
}
func (c Camera1) Stop() {
	fmt.Println("相机停止工作")
}

type Computer1 struct {
}

func (c Computer1) Working1(usb Usb1) {
	usb.Stop()
	//如果usb指向一个phone的结构体变量，还要执行call方法
	v, ok := usb.(Phone1)
	if ok == true {
		v.call()
	} else {
		fmt.Println("不是phone 所以还用特有的命令")
	}
	usb.Start()
}

func main() {
	//定义一个usb接口数组，可以存放phone和camera的结构体变量
	//如果不用接口 绝对不可能给数组放不同的结构体，所以这就是体现了多态的特征
	var usbArr [3]Usb1
	usbArr[0] = Phone1{"xiaomi"}
	usbArr[1] = Phone1{"vivo"}
	usbArr[2] = Camera1{"niko"}
	fmt.Println(usbArr)

	var computer Computer1
	for _, v := range usbArr {
		computer.Working1(v)
	}

}
