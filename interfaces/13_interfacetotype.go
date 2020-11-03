package main

//接口转换为其他类型，本例实现接口转换为其他类型
import "fmt"

//定义猪类
type pig1 struct {
}

//行走动物接口
type Walker1 interface {
	Walker()
}

func (p *pig1) Walker() {
	fmt.Println("pig1:walk")
}

func main() {
	p1 := new(pig1)
	var a Walker1 = p1
	p2 := a.(*pig1)
	//p2:=a.(*bird) 就会报错，因为a保存着的指针本来就是pig的，如果转换成bird就报错了
	fmt.Printf("p1=%p p2=%p", p1, p2) //两个变量指向同一个地址

}
