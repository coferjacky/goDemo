package cls1

import (
	"fmt"
	"packagetest/base"
)

//定义类1
type Class1 struct {
}

//实现class接口
func (c *Class1) Do() {
	fmt.Println("Class1")
}
func init() {
	base.Register("class1", func() base.Class {
		return new(Class1)
	})
}
