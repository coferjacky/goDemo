package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}
type Camera struct {
	Name string
}

func (c Phone) Start() {
	fmt.Println("手机开始开始工作")
}
func (c Phone) Stop() {
	fmt.Println("手机停止工作")
}
func (c Camera) Start() {
	fmt.Println("照相机开始开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Stop()
	usb.Start()
}

func main() {
	//定义一个usb接口数组，可以存放phone和camera的结构体变量
	//如果不用接口 绝对不可能给数组放不同的结构体，所以这就是体现了多态的特征
	var usbArr [3]Usb
	usbArr[0] = Phone{"xiaomi"}
	usbArr[1] = Phone{"vivo"}
	usbArr[2] = Camera{"niko"}
	fmt.Println(usbArr)
}
