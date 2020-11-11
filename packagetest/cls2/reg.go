package cls2

import (
	"fmt"
	"packagetest/base"
)

//定义类1
type Class2 struct {
}

//实现class接口
func (c *Class2) Do() {
	fmt.Println("Class2")
}
func init() {
	base.Register("class2", func() base.Class {
		return new(Class2)
	})
}
