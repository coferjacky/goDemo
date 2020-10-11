package main

import "fmt"

//指针形接收器
type Property struct {
	value int
}

//设置属性值
func (p *Property) Set_Value(v int) {
	p.value = v
}

//取属性值
func (p *Property) Value() int {
	return p.value
}
func main() {
	p := new(Property)
	p.Set_Value(100)
	fmt.Println(p.Value())
}
