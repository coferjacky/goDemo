package main

import (
	"packagetest/base"
	_ "packagetest/cls1"
	_ "packagetest/cls2"
)

func main() {
	c1 := base.Create("class1")
	c1.Do()
	c2 := base.Create("class2")
	c2.Do()
}
